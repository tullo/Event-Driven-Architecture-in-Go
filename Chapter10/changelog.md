# changelog

Testing strategy:
- [Unit](#unit-tests) tests
- [Integration](#integration-tests) tests
- Contract tests
- End-to-end tests

At each level of testing, we use the term system under test (SUT) to describe the component or components being tested.

## Unit tests

The bulk of the testing efforts.

These tests should be free of any dependencies, especially any I/O
- make use of test doubles
  - mocks (respond in certain ways to specific inputs)
  - spies (observable proxy of the real implementation)
  - stubs (responds with static or predictable responses)
  - fakes (same functionality as the real dependency) e.g. in-memory impl. of a repo

Test doubles are tools we can use to isolate the system or code under test from the rest of the system around it.

- Interaction with the dependency is NOT important => fakes & stubs
- Interaction with the dependency IS important => mocks & spies

> SUT should be very small; individual functions and methods.

```go
// https://dave.cheney.net/2019/05/07/prefer-table-driven-tests

tests := map[string]struct {
	fields  fields
	on      func(a *es.MockAggregate)
	want    ddd.Event
	wantErr bool
}{...}
for name, tt := range tests {
    t.Run(name, func(t *testing.T) {
        // arrange, act, and assert
        // aka Given-When-Then
    })
}
```

Mockery provides the ability to easily generate mocks for Golang interfaces using the stretchr/testify/mock
- https://github.com/stretchr/testify
- https://github.com/vektra/mockery

```go
//go:generate mockery --all --inpackage --case underscore

// --inpackage         create the mocks in the current package
// --case underscore   create the mocks using underscores in the filename
// --all               generate a mock for each interface that is found in dir and sub-dirs

// use of generated mocks
type mocks struct {
    baskets   *domain.MockBasketRepository
    stores    *domain.MockStoreRepository
    products  *domain.MockProductRepository
    publisher *ddd.MockEventPublisher[ddd.Event]
}

tests := map[string]struct {
    ...
    on      func(f mocks) // a function that accepts the mocks struct
    wantErr bool
}{...}

```

## Integration tests

Focus on testing the interactions between two components.

> The SUT will be the two components with any additional dependencies replaced with mocks.

- **Docker Compose**
  - `docker_compose`
  - https://golang.testcontainers.org/features/docker_compose/
- **LocalStack** - Develop and test your cloud and serverless apps offline!
  - https://localstack.cloud/
  - https://localstack.cloud/blog/2022-07-13-announcing-localstack-v1-general-availability/
- **Testcontainers**
  - https://golang.testcontainers.org/
  - start up a container or compose an environment that is controlled by code

```go
s.container, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
    ContainerRequest: testcontainers.ContainerRequest{
        Image:        "postgres:15.2-alpine",
        Hostname:     "postgres",
        ExposedPorts: []string{"5432/tcp"},
        Env: map[string]string{
            "POSTGRES_PASSWORD": "itsasecret",
        },
        Mounts: []testcontainers.ContainerMount{
            testcontainers.BindMount(initDir, "/docker-entrypoint-initdb.d"),
        },
        WaitingFor: wait.ForSQL("5432/tcp", "pgx", func(host string, port nat.Port) string {
            return fmt.Sprintf(dbUrl, host, port.Port())
        }),
    },
})
```

### Breaking tests into groups

- Running specific directories, files, or tests
- Go build constraints
- Using the short test option

All three options used together.:
- `go test ./internal/postgres -tags integration -short`

## Contract tests

We can use contract tests built by consumers’ expectations of an API or message
to verify whether a producer has implemented those expectations.

Contract tests are expected to run just as fast as unit tests since they do not deal with any real dependencies or test any logic.

> The SUT for a contract will be either the consumer and its expectation,
> or the producer and its API or message verification.

These tests are not just for testing between microservices – they can also be used to test your UI with its
backend API.

Contract testing allows us to focus on the APIs between providers and consumers, just like the integration tests do,
but it allows us to run the tests in isolation, similar to a unit test.

Contract testing comes in two forms:
- **Consumer**-driven contract testing (CDCT)
  - consumers **expectations**
  - mocked provider => contract
- **Provider**-driven contract testing (PDCT)
  - provider **verifications**
  - mocked consumer => contract


When a consumer’s contract is verified,
this can be shared with the consumer so that they know it will be OK to deploy with its API usage.

Likewise, a provider, having passed all of the contract verifications it was presented with,
will have the confidence in knowing it too can be deployed without any issues.

Consumers can make mistakes and have incorrect expectations, which could mean there is room to improve or add API documentation.

Providers may make a breaking change and will need to cooperate with the affected consumers to coordinate updates and releases.

- https://pact.io/ - libraries for many languages
- Pact Broker https://docs.pact.io/pact_broker - to share contracts
  - installed locally using a docker image
  - hosted with https://pactflow.io/
- CLI tools
  - rust: https://github.com/pact-foundation/pact-reference/tree/master/rust/pact_verifier_cli
    - new: can test both HTTP and Message Based interactions
  - docker: https://hub.docker.com/r/pactfoundation/pact-cli
  - executable: https://github.com/pact-foundation/pact_broker-client/releases

Pact Go workshop
- https://github.com/pact-foundation/pact-workshop-go

Testing Microservices - Contract Tests
- https://softwaremill.com/testing-microservices-contract-tests/

## E2E - End-to-end tests

E2E tests are used to test the expected functionality of the whole application.

E2E tests are often extensive and slow.

E2E testing will encompass the entire application,
including third-party services,
and have nothing replaced with any test doubles.

Record them in our feature test files using Gherkin https://cucumber.io/

```feature
Feature: Create Store

  As a store owner
  I should be able to create new stores

  Scenario: Creating a store called "Waldorf Books"
    Given a valid store owner
    And no store called "Waldorf Books" exists
    When I create the store called "Waldorf Books"
    Then a store called "Waldorf Books" exists
```

To make a feature file an **executable specification**, we will use the `godog` library.

```go
func (c *storesContext) register(ctx *godog.ScenarioContext) {
	ctx.Step(`^a valid store owner$`, c.aValidStoreOwner)
	ctx.Step(`^I create the store called "([^"]*)"$`, c.iCreateTheStoreCalled)
	ctx.Step(`^(?:ensure |expect )?a store called "([^"]*)" (?:to )?exists?$`, c.expectAStoreCalledToExist)
	ctx.Step(`^(?:ensure |expect )?no store called "([^"]*)" (?:to )?exists?$`, c.expectNoStoreCalledToExist)
}
```

We should not try to write features covering everything that the application does or can do.

Start with the critical flows to the business and then go from there.

Some flows may not automate very well and should be left for the testers to run through manually.
