# Load .env file
include .env
export

# ===============================
# CONFIG
# ===============================

MIGRATIONS_PATH=./cmd/migrate/migrations

# Build DB connection string from .env
DB_ADDR=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB)?sslmode=disable

# ===============================
# COMMANDS
# ===============================

.PHONY: migrate-create
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) down $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-force
version ?=
migrate-force:
	@if [ -z "$(version)" ]; then \
		echo "❌ Please provide version. Example: make migrate-force version=2"; \
		exit 1; \
	fi
	@echo "⚠️ Forcing migration version to $(version)"
	@migrate -path=$(MIGRATIONS_PATH) -database="$(DB_ADDR)" force $(version)

.PHONY: seed
seed:
	@go run cmd/migrate/seed/main.go


.PHONY: gen-docs
gen-docs:
	@swag init -g ./api/main.go -d cmd,internal && swag fmt

%:
	@:
