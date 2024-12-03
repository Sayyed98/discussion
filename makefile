# Variables
DB_URL=mysql://root:password@tcp(127.0.0.1:3306)/discussion
MIGRATE=migrate
MIGRATION_PATH=db/migration

# Targets
.PHONY: migrate-up migrate-down create-migration

# Apply all migrations
migrate-up:
	$(MIGRATE) -database "$(DB_URL)" -path $(MIGRATION_PATH) up

# Rollback the last migration
migrate-down:
	$(MIGRATE) -database "$(DB_URL)" -path $(MIGRATION_PATH) down 1

# Create a new migration file
create-migration:
	@if [ -z "$(name)" ]; then \
		echo "Error: Please provide a name for the migration (e.g., 'make create-migration name=init_schema')"; \
		exit 1; \
	fi
	$(MIGRATE) create -ext sql -dir $(MIGRATION_PATH) -seq $(name)