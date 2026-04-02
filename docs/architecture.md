# Architecture

## Overview

This project follows a simplified Hexagonal Architecture (Ports & Adapters), focused on clarity and testability. The goal is to keep a clean separation between business logic and external systems without over-engineering.

---

## Structure

```text
cmd/    → entrypoint (dependency wiring)
core/   → business logic (pure)
infra/  → external implementations
```

```text
core/
├── entities/
├── errors/
├── ports/
└── usecases/

infra/
├── blockchain/
├── db/
└── http/
```

---

## Core Principles

- Core is independent (no external dependencies)
- Ports define contracts (interfaces)
- Infra implements ports
- cmd wires everything together

```text
core → no dependencies
infra → depends on core
cmd → depends on all
```

---

## Use Cases

- set_storage_on_blockchain_usecase.go → write value to blockchain
- get_storage_from_blockchain_usecase.go → read value from blockchain
- sync_storage_to_database_usecase.go → sync blockchain value to database
- check_storage_consistency_usecase.go → compare database vs blockchain

---

## Data Flow

```text
HTTP → Use Case → Ports → Adapters (Blockchain / DB)
```

---

## Key Decisions

- Blockchain is the source of truth
- Database acts as a cache
- Minimal layering to avoid over-engineering

---

## Summary

This architecture keeps the core isolated, the infrastructure replaceable, and the system easy to test and evolve, while staying simple and practical.
