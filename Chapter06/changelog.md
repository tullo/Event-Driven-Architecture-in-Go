# changelog

- Using NATS Jetstream
- Removed aggregateEvent struct (postgres)
- Added UML diagrams (plantuml)
- Added env var: `NATS_URL=nats:4222`
- Added handlers.RegisterIntegrationEventHandlers(...)
- Added new jetstream package
  ```go
    NatsConfig struct {
	  URL    string `required:"true"`
	  Stream string `default:"mallbots"`
	}

    package jetstream
    
    type Stream struct {
	    streamName string
	    js         nats.JetStreamContext
    }

    func (s *Stream) Publish(ctx context.Context, topicName string, rawMsg am.RawMessage) (err error) {...}

    func (s *Stream) Subscribe(topicName string, handler am.MessageHandler[am.RawMessage], options ...am.SubscriberOption) error {...}
  ```
- Added application domain event handlers interface
    ```go
    type DomainEventHandlers interface {
	    OnCustomerRegistered(ctx context.Context, event ddd.Event) error
	    OnCustomerAuthorized(ctx context.Context, event ddd.Event) error
	    OnCustomerEnabled(ctx context.Context, event ddd.Event) error
	    OnCustomerDisabled(ctx context.Context, event ddd.Event) error
    }
    ```
- Added wait for stream method
    ```go
    func (a *app) waitForStream(ctx context.Context) error {
	    closed := make(chan struct{})
	    a.nc.SetClosedHandler(func(*nats.Conn) {
		close(closed)
	    })
	    group, gCtx := errgroup.WithContext(ctx)
	    group.Go(func() error {
		    fmt.Println("message stream started")
		    defer fmt.Println("message stream stopped")
		    <-closed
		    return nil
	    })
	    group.Go(func() error {
		    <-gCtx.Done()
		    return a.nc.Drain()
	    })
	    return group.Wait()
    }
    ```

## Messages / Events
- An event is a message
- A message is not always an event
- A message is a container with a payload
- Kinds of **payloads** in use:
  - `Integration event`: A state change that is communicated outside of its bounded context
  - `Command`: A request to perform work
  - `Query`: A request for some information
  - `Reply`: An informational response to either a command or query
- Kinds of **events** in use:
  - `Domain events`: Exist for the **shortest time**, never leave the application, do not require versioning, and are typically handled synchronously. These events are used to inform other application components about changes made to or by an aggregate.
  - `Event-sourced events`: Exist for the **longest time**, never leave the service boundary, require versioning, and are **handled synchronously**. These events keep a record of every change in state that is made to an aggregate.
  - `Integration events`: Exist for an **unknown amount of time**, are used by an unknown number of consumers, require versioning, and are typically handled asynchronously. These events are used to supply other components of the system with information regarding significant decisions or changes.
    - `notification events`
    - `event-carried state transfer events`

A notification event is going to be the smallest event you can send.

## Eventual consistency

`read-after-write` **inconsistency**, has to do with not being able to read the state change or new data immediately after writing it.

## Message-delivery guarantees

Available depending on the broker or libraries you use:
- `At-most-once message delivery`
  - The `Producer` **does not wait for an acknowledgment** from the Message Broker when it publishes a message
  - Message deduplication and idempotency are not a concern.
  - The `broker` **does not wait for any acknowledgment** from the `Consumer` before it deletes the message
  - If the `Consumer` **fails** to process the message, then the **message will be lost**
  - use cases, e.g:
    - collection of logs
    - processing messages from IoT devices
- `At-least-once message delivery`
  - `Producer` is **guaranteed** to have published the message to the `Message Broker`
  - the broker will **keep delivering** the message to the `Consumer` until the `Message Broker` has **received an acknowledgment** that the message has been received
  - A `Consumer` may receive the message **more than once**
    - must be utilizing message deduplication
    - must have implemented other idempotency measures
- `Axactly-once message delivery`
  - would be ideal for just about any situation
  - but it is extremely hard or downright impossible to achieve in most cases

## NATS JetStream
NATS (https://nats.io) is a messaging broker that supports subject-based
messaging and publish-subscribe (pub-sub).

An easy-to-use API and message model:
- **Subject**: A string
- **Payload**: A byte slice (max 64 MB)
- **Headers**: A map of string slices

Features that JetStream adds to NATS Core:
- **Message deduplication**: This can deduplicate messages that have been published more than once
- **Message replay**: Consumers may
  - receive **all messages**, or
  - receive messages after a **specific point** in the stream or
  - receive messages after a **specific timestamp**
- **Additional retention policies**: We can choose to keep messages if consumers exist with subscriptions to them or assign limits on the number of messages or total size of the stream

JetStream provides two components, the `Stream` and the `Consumer`.
- **Stream**: This is responsible for storing published messages for several 
`subjects`.
- **Consumer**: This is created as a `view` on the message store.
  - has a `cursor` that is used to iterate over the messages in a stream

---

## Added `am` package (asynchronous messaging)

- The `MessagePublisher` is able to publish anything into a stream
- The `MessageSubscriber` is returning a Message type of some kind
- `MessageStream` is a composition of MessagePublisher and MessageSubscriber interfaces that allow us to create a stream that will let us publish an `Event`
type and receive an `EventMessage` type

```go
type (
	Message interface {
		ddd.IDer
		MessageName() string
		Ack() error
		NAck() error
		Extend() error
		Kill() error
	}

	MessageHandler[O Message] interface {
		HandleMessage(ctx context.Context, msg O) error
	}

	MessageHandlerFunc[O Message] func(ctx context.Context, msg O) error

	MessagePublisher[I any] interface {
		Publish(ctx context.Context, topicName string, v I) error
	}

	MessageSubscriber[O Message] interface {
	  Subscribe(topicName string, handler MessageHandler[O], options ...SubscriberOption) error
	}

	MessageStream[I any, O Message] interface {
		MessagePublisher[I]
		MessageSubscriber[O]
	}
)
```

The `RawMessage` supports the need to serialize and deserialize events into something we can then pass into NATS JetStream using a non-vendor specific format.

```go
type (
	RawMessage interface {
		Message
		Data() []byte
	}

	rawMessage struct {
		id   string
		name string
		data []byte
	}
)

func (m rawMessage) ID() string          { return m.id }
func (m rawMessage) MessageName() string { return m.name }
func (m rawMessage) Data() []byte        { return m.data }
func (m rawMessage) Ack() error          { return nil }
func (m rawMessage) NAck() error         { return nil }
func (m rawMessage) Extend() error       { return nil }
func (m rawMessage) Kill() error         { return nil }
```

Implementation of the EventStream interface.

We have
- a `Publish()` method that accepts only the `ddd.Event` type
- a `Subscribe()` method that only accepts handlers that operate on `EventMessages`
- a registry to process the event payloads
- a stream that handles the RawMessage type for both the published input and the subscribed output types

```go
type (
	EventMessage interface {
		Message
		ddd.Event
	}

	EventPublisher  = MessagePublisher[ddd.Event]
	EventSubscriber = MessageSubscriber[EventMessage]
	EventStream     = MessageStream[ddd.Event, EventMessage]

	eventStream struct {
		reg    registry.Registry
		stream MessageStream[RawMessage, RawMessage]
	}
)

func (s eventStream) Publish(ctx context.Context, topicName string, event ddd.Event) error {...}

func (s eventStream) Subscribe(topicName string, handler MessageHandler[EventMessage], options ...SubscriberOption) error {...}
```

A protocol buffer message is used as the data container that is then used as the data for the raw message.

```protobuf
syntax = "proto3";

import "google/protobuf/timestamp.proto";
import "google/protobuf/struct.proto";

message EventMessageData {
  bytes payload = 1;
  google.protobuf.Timestamp occurred_at = 2;
  google.protobuf.Struct metadata = 3;
}
```
---

## Added `jetstream` package (NATS-specific functionality)

Adds infrastructure-specific implementations for NATS JetStream.

```protobuf
syntax = "proto3";

package jetstream;

message StreamMessage {
  string id = 1;
  string name = 2;
  bytes data = 3;
}
```

```go
type rawMessage struct {
	id       string
	name     string
	data     []byte
	acked    bool
	ackFn    func() error
	nackFn   func() error
	extendFn func() error
	killFn   func() error
}

func (m rawMessage)  ID() string          { return m.id }
func (m rawMessage)  MessageName() string { return m.name }
func (m rawMessage)  Data() []byte        { return m.data }
func (m *rawMessage) Ack() error          {...}
func (m *rawMessage) NAck() error         {...}
func (m rawMessage)  Extend() error       {...}
func (m *rawMessage) Kill() error         {...}

// Stream =====================================================================
// - Publish() is going to serialize the RawMessage into a NATS message.
// - Subscribe() is doing the opposite.
method is used when you want to create a subscription with competing consumers.
type Stream struct {
	streamName string
	js         nats.JetStreamContext
}

func (s *Stream) Publish(ctx context.Context, topicName string, rawMsg am.RawMessage) (err error) {...}

func (s *Stream) Subscribe(topicName string, handler am.MessageHandler[am.RawMessage], options ...am.SubscriberOption) error {...}

// Available options: JetStream Subscribe() or QueueSubscribe() method.
// The QueueSubscribe() method is to create a subscription with competing consumers.
```

---

## Public Events

A proto file is used to help with organizing `integration events` and to keep them
separate from the gRPC API messages.

```protobuf
// events.proto
syntax = "proto3";

package storespb;

message StoreCreated {
  string id = 1;
  string name = 2;
  string location = 3;
}

message StoreParticipationToggled {
  string id = 1;
  bool participating = 2;
}

message StoreRebranded {
  string id = 1;
  string name = 2;
}

message ProductAdded {
  string id = 1;
  string store_id = 2;
  string name = 3;
  string description = 4;
  string sku = 5;
  double price = 6;
}

message ProductRebranded {
  string id = 1;
  string name = 2;
  string description = 3;
}

message ProductPriceChanged {
  string id = 1;
  double delta = 2;
}

message ProductRemoved {
  string id = 1;
}
```

Publishing using constants and payload types available to the entire application.

```go
package storespb

const (
	StoreAggregateChannel = "mallbots.stores.events.Store"

	StoreCreatedEvent              = "storesapi.StoreCreated"
	StoreParticipatingToggledEvent = "storesapi.StoreParticipatingToggled"
	StoreRebrandedEvent            = "storesapi.StoreRebranded"

	ProductAggregateChannel = "mallbots.stores.events.Product"

	ProductAddedEvent          = "storesapi.ProductAdded"
	ProductRebrandedEvent      = "storesapi.ProductRebranded"
	ProductPriceIncreasedEvent = "storesapi.ProductPriceIncreased"
	ProductPriceDecreasedEvent = "storesapi.ProductPriceDecreased"
	ProductRemovedEvent        = "storesapi.ProductRemoved"
)




// baskets/... ================================================================

// StoreHandlers - HandleEvent()
func (h StoreHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case storespb.StoreCreatedEvent:
		return h.onStoreCreated(ctx, event)
	case storespb.StoreParticipatingToggledEvent:
		return h.onStoreParticipationToggled(ctx, event)
	case storespb.StoreRebrandedEvent:
		return h.onStoreRebranded(ctx, event)
	}

	return nil
}

// ProductHandlers - HandleEvent()
func (h ProductHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case storespb.ProductAddedEvent:
		return h.onProductAdded(ctx, event)
	case storespb.ProductRebrandedEvent:
		return h.onProductRebranded(ctx, event)
	case storespb.ProductPriceIncreasedEvent, storespb.ProductPriceDecreasedEvent:
		return h.onProductPriceChanged(ctx, event)
	case storespb.ProductRemovedEvent:
		return h.onProductRemoved(ctx, event)
	}

	return nil
}

// RegisterProductHandlers - Subscribe()
func RegisterProductHandlers(productHandlers ddd.EventHandler[ddd.Event], stream am.EventSubscriber) error {
	evtMsgHandler := am.MessageHandlerFunc[am.EventMessage](func(ctx context.Context, eventMsg am.EventMessage) error {
		return productHandlers.HandleEvent(ctx, eventMsg)
	})

	return stream.Subscribe(storespb.ProductAggregateChannel, evtMsgHandler, am.MessageFilter{
		storespb.ProductAddedEvent,
		storespb.ProductRebrandedEvent,
		storespb.ProductPriceIncreasedEvent,
		storespb.ProductPriceDecreasedEvent,
		storespb.ProductRemovedEvent,
	})
}

// RegisterStoreHandlers - Subscribe()
func RegisterStoreHandlers(storeHandlers ddd.EventHandler[ddd.Event], stream am.EventSubscriber) error {
	evtMsgHandler := am.MessageHandlerFunc[am.EventMessage](func(ctx context.Context, eventMsg am.EventMessage) error {
		return storeHandlers.HandleEvent(ctx, eventMsg)
	})

	return stream.Subscribe(storespb.StoreAggregateChannel, evtMsgHandler, am.MessageFilter{
		storespb.StoreCreatedEvent,
		storespb.StoreParticipatingToggledEvent,
		storespb.StoreRebrandedEvent,
	})
}

// ProductHandlers
// - onProductAdded()
// - onProductRebranded()
// - onProductPriceChanged()
// - onProductRemoved()
func (h ProductHandlers[T]) onProductAdded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storespb.ProductAdded)
	h.logger.Debug().Msgf(`ID: %s, Name: "%s", Price: "%d"`, payload.GetId(), payload.GetName(), payload.GetPrice())
	return nil
}

func (h ProductHandlers[T]) onProductRebranded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storespb.ProductRebranded)
	h.logger.Debug().Msgf(`ID: %s, Name: "%s", Description: "%s"`, payload.GetId(), payload.GetName(), payload.GetDescription())
	return nil
}

func (h ProductHandlers[T]) onProductPriceChanged(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storespb.ProductPriceChanged)
	h.logger.Debug().Msgf(`ID: %s, Delta: "%d"`, payload.GetId(), payload.GetDelta())
	return nil
}

func (h ProductHandlers[T]) onProductRemoved(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storespb.ProductRemoved)
	h.logger.Debug().Msgf(`ID: %s, Price: "%d"`, payload.GetId())
	return nil
}



// stores/... =================================================================

// IntegrationEventHandlers - Publish()
func (h IntegrationEventHandlers[T]) onStoreCreated(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domain.StoreCreated)
	return h.publisher.Publish(ctx, storespb.StoreAggregateChannel,
		ddd.NewEvent(storespb.StoreCreatedEvent, &storespb.StoreCreated{
			Id:       event.ID(),
			Name:     payload.Name,
			Location: payload.Location,
		}),
	)
}

func (h IntegrationEventHandlers[T]) onStoreParticipationEnabled(ctx context.Context, event ddd.AggregateEvent) error {
	return h.publisher.Publish(ctx, storespb.StoreAggregateChannel,
		ddd.NewEvent(storespb.StoreParticipatingToggledEvent, &storespb.StoreParticipationToggled{
			Id:            event.ID(),
			Participating: true,
		}),
	)
}

func (h IntegrationEventHandlers[T]) onStoreParticipationDisabled(ctx context.Context, event ddd.AggregateEvent) error {
	return h.publisher.Publish(ctx, storespb.StoreAggregateChannel,
		ddd.NewEvent(storespb.StoreParticipatingToggledEvent, &storespb.StoreParticipationToggled{
			Id:            event.ID(),
			Participating: false,
		}),
	)
}

func (h IntegrationEventHandlers[T]) onStoreRebranded(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domain.StoreRebranded)
	return h.publisher.Publish(ctx, storespb.StoreAggregateChannel,
		ddd.NewEvent(storespb.StoreRebrandedEvent, &storespb.StoreRebranded{
			Id:   event.ID(),
			Name: payload.Name,
		}),
	)
}
```

Swagger UI:
- http://localhost:8080/?urls.primaryName=Store%20Management#/Store/createStore

Possible resulting log messages for domain & integration events:

```log
12:50:41.505PM INF --> Stores.CreateStore
12:50:41.523PM INF --> Stores.Mall.On(stores.StoreCreated)
12:50:41.530PM INF <-- Stores.Mall.On(stores.StoreCreated)
12:50:41.530PM INF --> Stores.IntegrationEvents.On(stores.StoreCreated)
12:50:41.532PM INF <-- Stores.IntegrationEvents.On(stores.StoreCreated)
12:50:41.532PM INF <-- Stores.CreateStore
12:50:41.534PM INF --> Baskets.Store.On(storesapi.StoreCreated)
12:50:41.535PM DBG ID: 6148ebb8-d464-4760-8cc4-5dd36aab0164, Name: "Silvan", Location: "Lygten"
12:50:41.536PM INF <-- Baskets.Store.On(storesapi.StoreCreated)
```
