-include .env

.PHONY: build tidy run debug migrate-up migrate-down create-migration migrate-force migrate-status docker-up docker-down docker-rebuild

DB_URL=postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable
MIGRATE := migrate -path ./database/migrations -database $(DB_URL)
MIGRATE_NAME := $(if $(name),$(name),default_name)

build:
	@mkdir -p build
	@go build -o build/main main.go

tidy:
	@go mod tidy
	@go mod vendor

run:
	@GIN_MODE=release go run main.go

debug:
	@GIN_MODE=debug go run main.go

migrate-up:
	@$(MIGRATE) up
	@echo "Migrations applied successfully."

migrate-down:
	@$(MIGRATE) down
	@echo "Migrations rolled back successfully."

create-migration:
	@$(if $(name),,echo "Error: migration name required"; exit 1)
	@migrate create -ext sql -dir ./database/migrations -seq $(MIGRATE_NAME)
	@echo "Migration '$(MIGRATE_NAME)' created successfully."

migrate-force:
	@$(MIGRATE) force $(version)
	@echo "Migration forced successfully."

migrate-status:
	@$(MIGRATE) version

docker-up:
	@docker compose up -d --build

docker-down:
	@docker compose down

docker-rebuild:
	@docker compose down
	@docker compose up -d --build