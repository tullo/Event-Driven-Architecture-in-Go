# changelog

```log
INF --> Baskets.CheckoutBasket
INF --> Baskets.DomainEvents.On(baskets.BasketCheckedOut)
INF <-- Baskets.DomainEvents.On(baskets.BasketCheckedOut)
INF <-- Baskets.CheckoutBasket
INF --> Ordering.IntegrationEvents.On(basketsapi.BasketCheckedOut)
INF --> Ordering.CreateOrder
INF --> Ordering.DomainEvents.On(ordering.OrderCreated)
INF <-- Ordering.DomainEvents.On(ordering.OrderCreated)
INF <-- Ordering.CreateOrder
INF <-- Ordering.IntegrationEvents.On(basketsapi.BasketCheckedOut)
INF --> Notifications.IntegrationEvents.On(ordersapi.OrderCreated)
INF --> Notifications.NotifyOrderCreated
INF <-- Notifications.NotifyOrderCreated
INF <-- Notifications.IntegrationEvents.On(ordersapi.OrderCreated)
INF --> COSEC.IntegrationEvents.On(ordersapi.OrderCreated)
INF <--COSEC.IntegrationEvents.On(ordersapi.OrderCreated)
INF --> Customers.Commands.On(customersapi.AuthorizeCustomer)
INF --> Customers.AuthorizeCustomer
INF <-- Customers.AuthorizeCustomer
INF <-- Customers.Commands.On(customersapi.AuthorizeCustomer)
INF --> COSEC.CreateOrderSaga.On(am.Success)
INF <-- COSEC.CreateOrderSaga.On(am.Success)
INF --> Depot.Commands.On(depotapi.CreateShoppingListCommand)
INF --> Depot.CreateShoppingList
INF <-- Depot.CreateShoppingList
INF <-- Depot.Commands.On(depotapi.CreateShoppingListCommand)
INF --> COSEC.CreateOrderSaga.On(depotapi.CreatedShoppingListReply)
INF <-- COSEC.CreateOrderSaga.On(depotapi.CreatedShoppingListReply)
INF --> Payments.Commands.On(paymentsapi.ConfirmPayment)
INF --> Payments.ConfirmPayment
INF <-- Payments.ConfirmPayment
INF <-- Payments.Commands.On(paymentsapi.ConfirmPayment)
INF --> COSEC.CreateOrderSaga.On(am.Success)
INF <-- COSEC.CreateOrderSaga.On(am.Success)
INF --> Depot.Commands.On(depotapi.InitiateShoppingCommand)
INF <-- Depot.Commands.On(depotapi.InitiateShoppingCommand)
INF --> COSEC.CreateOrderSaga.On(am.Success)
INF <-- COSEC.CreateOrderSaga.On(am.Success)
INF --> Ordering.Commands.On(ordersapi.ApproveOrder)
INF --> Ordering.DomainEvents.On(ordering.OrderApproved)
INF <-- Ordering.DomainEvents.On(ordering.OrderApproved)
INF <-- Ordering.Commands.On(ordersapi.ApproveOrder)
INF --> COSEC.CreateOrderSaga.On(am.Success)
INF <-- COSEC.CreateOrderSaga.On(am.Success)
```

COSEC => Create-Order-Saga-Execution-Coordinator

## Addresses dual write problem

Create a transaction in PostgreSQL, into which we put our writes in order to combine them into a single write operation.

This is called the `Transactional Outbox` pattern.

Write the messages we publish into a database alongside all of other changes.

Changes to the modules
- Single database connection and transaction used for the lifetime of an entire request
- Intercepting and redirecting all outgoing messages into the database

```sql
-- === INBOX ==============================================

-- holds every incoming RawMessage instance
-- that the baskets module receives
-- used for deduplication and idempotent message delivery
  CREATE TABLE baskets.inbox
  (
    id          text NOT NULL,
    name        text NOT NULL,
    subject     text NOT NULL,
    data        bytea NOT NULL,
    received_at timestamptz NOT NULL,
    PRIMARY KEY (id)
  );
-- CREATE TABLE customers.inbox ...
-- CREATE TABLE depot.inbox ...
-- CREATE TABLE ordering.inbox ...
-- CREATE TABLE payments.inbox ...
-- CREATE TABLE search.inbox ...
-- CREATE TABLE stores.inbox ...
-- CREATE TABLE cosec.inbox ...

-- === OUTBOX =============================================
  CREATE TABLE baskets.outbox
  (
    id           text NOT NULL,
    name         text NOT NULL,
    subject      text NOT NULL,
    data         bytea NOT NULL,
    published_at timestamptz,
    PRIMARY KEY (id)
  );

-- CREATE TABLE customers.outbox ...
-- CREATE TABLE depot.outbox ...
-- CREATE TABLE ordering.outbox ...
-- CREATE TABLE payments.outbox ...
-- CREATE TABLE search.outbox ...
-- CREATE TABLE stores.outbox ...
-- CREATE TABLE cosec.outbox ...
```

Inbox middleware

```go
// evtHandlers implements RawMessageHandler
evtHandlers := am.RawMessageHandlerWithMiddleware(
	am.NewEventMessageHandler(
		di.Get(ctx, "registry").(registry.Registry),
		di.Get(ctx, "integrationEventHandlers").(ddd.EventHandler[ddd.Event]),
	),
	di.Get(ctx, "inboxMiddleware").(am.RawMessageHandlerMiddleware),
)
```

Outbox middleware to catch outgoing messages and save them into the outbox table

```go
// example: baskets.outbox
outboxStore := pg.NewOutboxStore("baskets.outbox", tx)
return am.RawMessageStreamWithMiddleware(
	c.Get("stream").(am.RawMessageStream),
	tm.NewOutboxStreamMiddleware(outboxStore),
)
```

## Added dependency injection (DI) package

Creates repository instances created with transactions.
- can create either singleton instances for the lifetime of
the application
- can create scoped instances that will exist only for the lifetime of a request

The `internal/di` package provides a container that will be used to register factory functions for singletons and scoped values

```go
type Container interface {
	AddSingleton(key string, fn DepFactoryFunc)
	AddScoped(key string, fn DepFactoryFunc)
	Scoped(ctx context.Context) context.Context
	Get(key string) any
}

// fetch values from the container inside the contexts
func Get(ctx context.Context, key string) any {
	ctn, ok := ctx.Value(containerKey).(*container)
	if !ok {
		panic("container does not exist on context")
	}

	return ctn.Get(key)
}
```

## Modules startup using DI container

Dependency injection for modules:
- baskets
- cosec
- customers
- depot
- ordering
- payments
- search
- stores

Done in Startup method of modules.

Example from baskets module - (simplified code snippet): 
```go
func (m *Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {

	container := di.New()

	// setup Driven adapters
	container.AddSingleton("registry", func(c di.Container) (any, error) {})
	container.AddSingleton("logger", func(c di.Container) (any, error) {})
	container.AddSingleton("stream", func(c di.Container) (any, error) {})
	container.AddSingleton("domainDispatcher", func(c di.Container) (any, error) {})
	container.AddSingleton("db", func(c di.Container) (any, error) {})
	container.AddSingleton("conn", func(c di.Container) (any, error) {})
	container.AddSingleton("outboxProcessor", func(c di.Container) (any, error) {})

	container.AddScoped("tx", func(c di.Container) (any, error) {})
	container.AddScoped("txStream", func(c di.Container) (any, error) {})
	container.AddScoped("eventStream", func(c di.Container) (any, error) {})
	container.AddScoped("inboxMiddleware", func(c di.Container) (any, error) {})
	container.AddScoped("aggregateStore", func(c di.Container) (any, error) {})
	container.AddScoped("baskets", func(c di.Container) (any, error) {})
	container.AddScoped("stores", func(c di.Container) (any, error) {})
	container.AddScoped("products", func(c di.Container) (any, error) {})

	// setup application
	container.AddScoped("app", func(c di.Container) (any, error) {})
	container.AddScoped("domainEventHandlers", func(c di.Container) (any, error) {})
	container.AddScoped("integrationEventHandlers", func(c di.Container) (any, error) {})

	// setup Driver adapters
	// using the container instead of any dependencies directly
	grpc.RegisterServerTx(container, mono.RPC())
	handlers.RegisterDomainEventHandlersTx(container)
	handlers.RegisterIntegrationEventHandlersTx(container)
	startOutboxProcessor(ctx, container)
}
```

Accessing instance from the container

```go
db1 := container.Get("db") // singleton
tx1 := container.Get("tx") // scoped

ctx := container.Scoped(context.Background())

db2 := di.Get(ctx, "db") // same instance as db1
tx2 := di.Get(ctx, "tx") // entirely new instance

tx3 := di.Get(ctx, "tx") // same instance as tx2

ctx = container.Scoped(context.TODO())
tx4 := di.Get(ctx, "tx") // entirely new instance



func (c *container) Get(key string) any {
	info, exists := c.deps[key]
	if info.scope == Singleton {
		return c.getFromParent(info)
	}

	return c.get(info)
}

func (c *container) Scoped(ctx context.Context) context.Context {
	// Use context Values only for request-scoped data
	// that transits processes and APIs.
	return context.WithValue(ctx, containerKey, c.scoped())
}
```

---

## The outbox message processor

The processor fetches a block of messages, publish each of them,
and then update the table to mark them as actually having been published.

The processor itself suffers from a dual write problem.

But we have deduplication in place, so the modules will be protected from any processor failures.

```go
func startOutboxProcessor(ctx context.Context, container di.Container) {
	outboxProcessor := container.Get("outboxProcessor").(tm.OutboxProcessor)
	logger := container.Get("logger").(zerolog.Logger)

	go func() {
		err := outboxProcessor.Start(ctx)
		if err != nil {
			logger.Error().Err(err).Msg("baskets outbox processor encountered an error")
		}
	}()
}
```

The processor will fetch up to 50 messages at a time to publish
and will wait for half a second in between queries
looking for messages that need to be published.

```go
const pollingInterval = 500 * time.Millisecond

func (p outboxProcessor) processMessages(ctx context.Context) error {
	timer := time.NewTimer(0)
	for {
		// fetch up to 50 messages at a time
		msgs, err := p.store.FindUnpublished(ctx, messageLimit)
		if err != nil {
			return err
		}

		if len(msgs) > 0 {
			ids := make([]string, len(msgs))
			// 1st write: publish messages
			for i, msg := range msgs {
				ids[i] = msg.ID()
				err = p.publisher.Publish(ctx, msg.Subject(), msg)
				if err != nil {
					return err
				}
			}
			// 2nd write: update database
			err = p.store.MarkPublished(ctx, ids...)
			if err != nil {
				return err
			}

			// no waiting: poll again immediately
			continue
		}

		if !timer.Stop() {
			select {
			case <-timer.C:
			default:
			}
		}

		// wait a short time before polling again
		timer.Reset(pollingInterval)

		select {
		case <-ctx.Done():
			return nil
		case <-timer.C:
		}
	}
}
```
