# changelog

- go 1.18 (`any` instead of empty interface)
- Updated the `Event` interface and declared a new `EventPayload` interface
- Added an event struct and a new Event constructor
- Replaced the `Aggregate` interface with a struct and added a constructor for it
- Embedded an updated `Entity` into `Aggregate` and added a constructor for it as well
- Updated the `AddEvent()` method to track Aggregate information on the events
- Updated `EventDispatcher` to use **generics** to avoid losing type safety or creating many
new versions
- Updated the modules to correctly build **new Aggregate instances with new constructors**
- Moved the **event names** into **constants** and used them in calls to AddEvent()
- Updated the handlers to perform **type assertions** on the event `Payload()`
- Added the internal event sourcing package `es`
- Switched to **event-sourced** aggregates for basket, order, product, store.
  - Implemented `ApplyEvent(event ddd.Event) error`
- Store changes
  ```go
  store.AddEvent(StoreCreatedEvent, &StoreCreated{
		Name:     name,
		Location: location,
	})
  ```
- Repositories and Stores
    ```go
    // domain
    domain.BasketRepository.Load
    domain.BasketRepository.Save

    domain.OrderRepository.Load
    domain.OrderRepository.Save

    domain.ProductRepository.Load
    domain.ProductRepository.Save

    domain.CatalogRepository.AddProduct
    domain.CatalogRepository.Rebrand
    domain.CatalogRepository.UpdatePrice
    domain.CatalogRepository.RemoveProduct
    domain.CatalogRepository.Find
    domain.CatalogRepository.GetCatalog

    domain.MallRepository.AddStore
    domain.MallRepository.SetStoreParticipation
    domain.MallRepository.RenameStore
    domain.MallRepository.Find
    domain.MallRepository.All
    domain.MallRepository.AllParticipating
    
    // database
    postgres.EventStore.Load
    postgres.EventStore.Save

    postgres.SnapshotStore.Load
    postgres.SnapshotStore.Save
    ```
