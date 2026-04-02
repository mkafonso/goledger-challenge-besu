# Go API (Besu + Postgres)

https://github.com/user-attachments/assets/889aa456-1a46-4963-b69a-f53ca4950c45

This folder contains the Go application that exposes an HTTP API to:

- write a value to the `SimpleStorage` smart contract deployed on the local Besu devnet
- read the current value from the blockchain
- sync the current blockchain value to Postgres
- compare (check) blockchain vs database values

More documentation is available in the repository-level [docs](../docs) folder.

## Run

If all prerequisites from the repository root README are installed, you can run the full project (Postgres + migrations + Besu devnet + contract deploy + API) with a single command from the repository root:

```bash
make devnet-deploy
```

Default API URL: `http://localhost:8080`

## Configuration

The API loads environment variables from `app/.env`.

- If `app/.env` does not exist, `make devnet-deploy` copies it from `app/.env.example`.
- `BLOCKCHAIN_PRIVATE_KEY` should match the pre-funded deployer key from `SimpleStorage/.env.example`.
- `BLOCKCHAIN_CONTRACT_ADDRESS` should match the deployed `SimpleStorage` contract address printed by `make devnet-deploy` (the example value works for a clean local devnet).

## Endpoints

- `POST /api/v1/storage` — sets a new value on-chain
- `GET /api/v1/storage` — reads the current on-chain value
- `POST /api/v1/storage/sync` — reads on-chain and persists to Postgres
- `GET /api/v1/storage/check` — returns `true` if DB value matches on-chain value

## Tests

Run all Go tests from the repository root:

```bash
make test
```
