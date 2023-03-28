# changelog

- added used event-carried state transfer to decouple modules
- added local caches for stores, product and customers
  - `CREATE TABLE baskets.stores_cache`
  - `CREATE TABLE baskets.products_cache`
  - `CREATE TABLE depot.stores_cache`
  - `CREATE TABLE depot.products_cache`
  - `CREATE TABLE search.customers_cache`
  - `CREATE TABLE search.stores_cache`
  - `CREATE TABLE search.products_cache`

Cache and Fallback repositories

```go
// baskets
stores := postgres.NewStoreCacheRepository("baskets.stores_cache", mono.DB(), grpc.NewStoreRepository(conn))
products := postgres.NewProductCacheRepository("baskets.products_cache", mono.DB(), grpc.NewProductRepository(conn))

// depots
stores := postgres.NewStoreCacheRepository("depot.stores_cache", mono.DB(), grpc.NewStoreRepository(conn))
products := postgres.NewProductCacheRepository("depot.products_cache", mono.DB(), grpc.NewProductRepository(conn))

// search
customers := postgres.NewCustomerCacheRepository("search.customers_cache", mono.DB(), grpc.NewCustomerRepository(conn))
stores := postgres.NewStoreCacheRepository("search.stores_cache", mono.DB(), grpc.NewStoreRepository(conn))
products := postgres.NewProductCacheRepository("search.products_cache", mono.DB(), grpc.NewProductRepository(conn))
```

Customer, Product and Store Cache Interfaces

```go
// baskets ====================================================================
type ProductCacheRepository interface {
	Add(ctx context.Context, productID, storeID, name string, price float64) error
	Rebrand(ctx context.Context, productID, name string) error
	UpdatePrice(ctx context.Context, productID string, delta float64) error
	Remove(ctx context.Context, productID string) error
	ProductRepository
}

type StoreCacheRepository interface {
	Add(ctx context.Context, storeID, name string) error
	Rename(ctx context.Context, storeID, name string) error
	StoreRepository
}

// depots =====================================================================
type ProductCacheRepository interface {
	Add(ctx context.Context, productID, storeID, name string) error
	Rebrand(ctx context.Context, productID, name string) error
	Remove(ctx context.Context, productID string) error
	ProductRepository
}

type StoreCacheRepository interface {
	Add(ctx context.Context, storeID, name, location string) error
	Rename(ctx context.Context, storeID, name string) error
	StoreRepository
}

// notification ===============================================================
type CustomerCacheRepository interface {
	Add(ctx context.Context, customerID, name, smsNumber string) error
	UpdateSmsNumber(ctx context.Context, customerID, smsNumber string) error
	CustomerRepository
}

// search =====================================================================
type CustomerCacheRepository interface {
	Add(ctx context.Context, customerID, name string) error
	CustomerRepository
}

type ProductCacheRepository interface {
	Add(ctx context.Context, productID, storeID, name string) error
	Rebrand(ctx context.Context, productID, name string) error
	Remove(ctx context.Context, productID string) error
	ProductRepository
}

type StoreCacheRepository interface {
	Add(ctx context.Context, storeID, name string) error
	Rename(ctx context.Context, storeID, name string) error
	StoreRepository
}
```

The receiving code is acting as an anti-corruption layer.
1. local cache lookup
2. fallback lookup
3. add it to local cache (ignoring pgerrcode.UniqueViolation)

---

- Store Management module is publishing.
- The Shopping Baskets module now consumes the events to create a local cache.

```log
INF --> Stores.CreateStore
INF --> Stores.Mall.On(stores.StoreCreated)
INF <-- Stores.Mall.On(stores.StoreCreated)
INF --> Stores.IntegrationEvents.On(stores.StoreCreated) <== asynchronously published
INF <-- Stores.IntegrationEvents.On(stores.StoreCreated)
INF <-- Stores.CreateStore
INF --> Depot.Store.On(storesapi.StoreCreated) <========== asynchronously consumed
INF --> Baskets.Store.On(storesapi.StoreCreated) <======== asynchronously consumed
INF --> Payments.Store.On(storesapi.StoreCreated) <======= asynchronously consumed
INF <-- Depot.Store.On(storesapi.StoreCreated)
INF <-- Baskets.Store.On(storesapi.StoreCreated)
INF <-- Payments.Store.On(storesapi.StoreCreated)
```

---


- Customers module is publishing.
- Payments module now consumes the events to create a local cache.

```log
INF --> Customers.RegisterCustomer
INF --> Customers.IntegrationEvents.On(customers.CustomerRegistered) <== asynchronously published
INF <-- Customers.IntegrationEvents.On(customers.CustomerRegistered)
INF <-- Customers.RegisterCustomer
INF --> Payments.Customer.On(customersapi.CustomerRegistered) <======= asynchronously consumed
INF --> Notifications.Customer.On(customersapi.CustomerRegistered) <== asynchronously consumed
INF <-- Payments.Customer.On(customersapi.CustomerRegistered)
INF <-- Notifications.Customer.On(customersapi.CustomerRegistered)
```

## Documenting the asynchronous API

AsyncAPI specification
- https://www.asyncapi.com/
- AsyncAPI generator tool

EventCatalog
- https://eventcatalog.dev
- uses Markdown files and functions
- The generated site can provide a visualization of the relationships that services have through their events.
- The site can even render a 3D node graph of the entire system with animations showing the direction in which state flows.

### Store Management AsyncAPI 1.0.0
- http://10.141.159.14:8080/stores-asyncapi/


## Read Model - Search

Uses data types for the columns that will have the fewest issues should the incoming data types change.

```sql
  CREATE TABLE search.orders
  (
    order_id       text NOT NULL,
    customer_id    text NOT NULL,
    customer_name  text NOT NULL,
    items          bytea NOT NULL,
    status         text NOT NULL,
    product_ids    text ARRAY NOT NULL,
    store_ids      text ARRAY NOT NULL,
    created_at     timestamptz NOT NULL DEFAULT NOW(),
    updated_at     timestamptz NOT NULL DEFAULT NOW(),
    PRIMARY KEY (order_id)
  );
```

When an order is created
- we combine it with the data already stored in the database
- creating our rich search model.

After the read model has been created
- it will receive additional updates as the status changes.

