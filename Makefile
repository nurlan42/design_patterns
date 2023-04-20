# Define variables for the migration command and its arguments
MIGRATE_CMD=migrate
MIGRATE_PATH=internal/migration
MIGRATE_DATABASE=postgresql://admin:secret@localhost:5432/todo?sslmode=disable
MIGRATE_VERBOSE=-verbose
MIGRATE_UP=up

# Define the target for running the migration command
migrate:
	$(MIGRATE_CMD) -path $(MIGRATE_PATH) -database $(MIGRATE_DATABASE) $(MIGRATE_VERBOSE) $(MIGRATE_UP)

schema:
	create -ext sql -dir internal/migration -seq init_schema

run:
	go run cmd/main.go
