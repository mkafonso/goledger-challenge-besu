# Checklist

## Setup

- [x] Fork repository
- [x] Create `/app` project structure

---

## Core Implementation

- [x] Define storage entity
- [x] Create ports:
  - [x] StorageBlockchainProvider (GetStorageFromBlockchain, SetStorageOnBlockchain)
  - [x] StorageRepository (GetStorage, SetStorage)
- [x] Implement use cases:
  - [x] SET set_storage_on_blockchain_usecase.go
  - [x] GET get_storage_from_blockchain_usecase.go
  - [x] SYNC sync_storage_to_database_usecase.go
  - [x] CHECK check_storage_consistency_usecase.go

---

## Business Rules Validation

- [x] SET writes only to blockchain
- [x] GET reads only from blockchain
- [x] SYNC persists blockchain value to DB
- [x] CHECK compares DB vs blockchain

---

## Testing

- [x] Unit tests for entities
- [x] Unit tests for usecases

---

## Error Handling

- [x] Define domain errors

---

## Documentation

- [x] Architecture documentation

---

## Final Validation

- [ ] Can set value on blockchain
- [ ] Can get value from blockchain
- [ ] Can sync value to DB
- [ ] Can compare DB vs blockchain
- [ ] API works end-to-end
- [ ] Tests pass
- [ ] Code is clean and organized
