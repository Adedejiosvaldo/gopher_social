include .envrc
MIGRATION_DIR = ./cmd/migrate/migrations


.PHONY: migration
migration:
	@echo "Creating migration..."
	@migrate create -seq -ext sql -dir $(MIGRATION_DIR) $(filter-out $@,$(MAKECMDGOALS))
	@echo "Migration created."


.PHONY: migrate-up
migrate-up:
	@echo "Running migrations..."
	@migrate -path=$(MIGRATION_DIR) -database=$(DATABASE_URL) up
	@echo "Migrations completed."

.PHONY: migrate-down
migrate-down:
	@echo "Running migrations..."
	@migrate -path=$(MIGRATION_DIR) -database=$(DATABASE_URL) down $(FILTER_OUT $@,$(MAKECMDGOALS))
	@echo "Migrations completed."

.PHONY: migrate-force
migrate-force:
	@echo "Forcing migration version..."
	@migrate -path=$(MIGRATION_DIR) -database=$(DATABASE_URL) force $(filter-out $@,$(MAKECMDGOALS))
	@echo "Migration forced."
