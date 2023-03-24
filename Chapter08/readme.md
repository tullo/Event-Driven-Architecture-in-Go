# changelog

Introduces several different options for performing complex operations
across different components in a distributed and asynchronous way.

Possible options to handle consistency across a distributed system
- `Two-Phase Commit (2PC)` - prepare and commit, **ACID guarantees**
- `Choreographed Saga` - does rely on individual components **publishing events**
- `Orchestrated Saga` - uses a saga **execution coordinator** to send **commands** to the components

A `saga` is a **sequence of steps** that define the `actions` and `compensating actions`
for the system components that are involved, aka the `saga participants`.

Sagas drop support for the isolation guarantee, making them **ACD transactions**.

Any **changes that are made will be visible to concurrent operations** while the other steps are being run.

A saga can be **long-lived**.

### Added support for the Command and Reply messages

When a Command message is handled, the outcome will be a Reply message.

```go
type (
	Reply interface {
		ID() string
		ReplyName() string
		Payload() ReplyPayload
		Metadata() Metadata
		OccurredAt() time.Time
	}

	Command interface {
		IDer
		CommandName() string
		Payload() CommandPayload
		Metadata() Metadata
		OccurredAt() time.Time
	}

	ReplyHandler[T Reply] interface {
		HandleReply(ctx context.Context, reply T) error
	}

	CommandHandler[T Command] interface {
		HandleCommand(ctx context.Context, cmd T) (Reply, error)
	}
)
```

### The Orchestrator

The primary job of our Orchestrator implementation is to handle the incoming replies so that it can determine which step to execute, as well as when to fail over and begin compensating.

```go
type (
	Orchestrator[T any] interface {
		Start(ctx context.Context, id string, data T) error
		ReplyTopic() string
		HandleReply(ctx context.Context, reply ddd.Reply) error
	}
)
```

### Saga definition

The Saga definition provides a single location for all of the logic on how the saga should operate.

```go
type (
	Saga[T any] interface {
		AddStep() SagaStep[T]
		Name() string
		ReplyTopic() string
		getSteps() []SagaStep[T]
	}
)
```
### Saga Steps

Steps are where all the logic of a Saga is contained.
They generate the Command messages that are sent to participants and can modify the data for the associated saga.

Each Step has, at a minimum, either an action or compensating action defined.

```go
type (
	SagaStep[T any] interface {
		Action(fn StepActionFunc[T]) SagaStep[T]
		Compensation(fn StepActionFunc[T]) SagaStep[T]
		OnActionReply(replyName string, fn StepReplyHandlerFunc[T]) SagaStep[T]
		OnCompensationReply(replyName string, fn StepReplyHandlerFunc[T]) SagaStep[T]
		isInvocable(compensating bool) bool
		execute(ctx context.Context, sagaCtx *SagaContext[T]) stepResult[T]
		handle(ctx context.Context, sagaCtx *SagaContext[T], reply ddd.Reply) error
	}
)
```

## Added module cosec

Short for Create-Order-Saga-Execution-Coordinator

```go
// cosec/internal
func NewCreateOrderSaga() sec.Saga[*models.CreateOrderData] {
	saga := createOrderSaga{
		Saga: sec.NewSaga[*models.CreateOrderData](CreateOrderSagaName, CreateOrderReplyChannel),
	}

	// 0. -RejectOrder
	saga.AddStep().
		Compensation(saga.rejectOrder)

	// 1. AuthorizeCustomer
	saga.AddStep().
		Action(saga.authorizeCustomer)

	// 2. CreateShoppingList, -CancelShoppingList
	saga.AddStep().
		Action(saga.createShoppingList).
		OnActionReply(depotpb.CreatedShoppingListReply, saga.onCreatedShoppingListReply).
		Compensation(saga.cancelShoppingList)

	// 3. ConfirmPayment
	saga.AddStep().
		Action(saga.confirmPayment)

	// 4. InitiateShopping
	saga.AddStep().
		Action(saga.initiateShopping)

	// 5. ApproveOrder
	saga.AddStep().
		Action(saga.approveOrder)

	return saga
}
```
