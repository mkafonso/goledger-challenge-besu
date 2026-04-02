BESU_DIR     := besu
CONTRACT_DIR := SimpleStorage
ENV_FILE     := $(CONTRACT_DIR)/.env
ENV_EXAMPLE  := $(CONTRACT_DIR)/.env.example

APP_DIR         := app
APP_ENV_FILE    := $(APP_DIR)/.env
APP_ENV_EXAMPLE := $(APP_DIR)/.env.example

DB_NAME      := app-besu-node-db
DB_USER      := postgres
DB_PASSWORD  := postgres
DB_PORT      := 5432
DB_CONTAINER := postgres-devnet

TOOLS_DIR := $(CURDIR)/bin
Tparse    := $(TOOLS_DIR)/tparse

.PHONY: devnet stop-devnet deploy devnet-deploy api postgres createdb migrateup generate_sqlc sqlc_down sqlc_up test tools

# ----------------------------
# TOOLS (LOCAL INSTALL)
# ----------------------------
tools:
	@mkdir -p $(TOOLS_DIR)
	@echo "Installing dev tools locally..."
	@GOBIN=$(TOOLS_DIR) go install github.com/mfridman/tparse@latest

# ----------------------------
# BESU
# ----------------------------
devnet:
	cd $(BESU_DIR) && ./startBesu.sh

stop-devnet:
	cd $(BESU_DIR) && ./stopBesu.sh


# ----------------------------
# POSTGRES
# ----------------------------
postgres:
	@echo "Starting Postgres container..."
	@docker run -d \
		--name $(DB_CONTAINER) \
		-e POSTGRES_USER=$(DB_USER) \
		-e POSTGRES_PASSWORD=$(DB_PASSWORD) \
		-e POSTGRES_DB=$(DB_NAME) \
		-p $(DB_PORT):5432 \
		postgres:16 || true


createdb:
	@echo "Ensuring database exists..."
	@docker exec -i $(DB_CONTAINER) psql -U $(DB_USER) -tc "SELECT 1 FROM pg_database WHERE datname='$(DB_NAME)'" | grep -q 1 || \
	docker exec -i $(DB_CONTAINER) psql -U $(DB_USER) -c "CREATE DATABASE $(DB_NAME)"


# ----------------------------
# MIGRATIONS
# ----------------------------
migrateup:
	@echo "Running migrations..."
	@bash app/script/ensure_migrate.sh
	@migrate -path app/infra/db/migrations \
		-database "postgres://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable" up


# ----------------------------
# SQLC
# ----------------------------
generate_sqlc:
	@echo "Generating sqlc..."
	@sqlc generate --file app/infra/db/sqlc.yaml

sqlc_up:
	@if command -v sqlc >/dev/null 2>&1; then \
		echo "Generating sqlc..."; \
		sqlc generate --file app/infra/db/sqlc.yaml; \
	else \
		echo "sqlc not found — skipping generation"; \
	fi

sqlc_down:
	@echo "Nothing to do"


# ----------------------------
# CONTRACT DEPLOY
# ----------------------------
deploy:
	@if [ ! -f $(ENV_FILE) ]; then \
		echo "No .env found — copying from .env.example"; \
		cp $(ENV_EXAMPLE) $(ENV_FILE); \
	fi
	@set -a && . ./$(ENV_FILE) && set +a && \
		cd $(CONTRACT_DIR) && forge script script/SimpleStorage.s.sol:SimpleStorageScript \
		--rpc-url besu \
		--broadcast


# ----------------------------
# FULL BOOTSTRAP
# ----------------------------
devnet-deploy:
	@if [ ! -f $(APP_ENV_FILE) ]; then \
		echo "No app .env found — copying from .env.example"; \
		cp $(APP_ENV_EXAMPLE) $(APP_ENV_FILE); \
	fi
	@make postgres
	@sleep 2
	@make createdb
	@make migrateup
	@make sqlc_up
	@make devnet
	@sleep 3
	@make deploy
	@make api


# ----------------------------
# API
# ----------------------------
api:
	cd app && go run ./cmd/api/main.go


# ----------------------------
# TEST (NORMAL)
# ----------------------------
test:
	@echo "Running tests (standard)..."
	@cd app && go test ./... -v


# ----------------------------
# TEST (BEAUTIFY)
# ----------------------------
test-beautify: tools
	@echo "Running tests (beautified)..."
	@cd app && go test ./... -json | $(Tparse) -all