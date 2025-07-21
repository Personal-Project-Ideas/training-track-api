# Load environment variables from .env
include .env
export

# Define variables
LIQUIBASE=liquibase
CHANGELOG_FILE=migrations/changelog.yml
MIGRATION_DIR=migrations

# Run API build
build:
	@echo "Building the Go API..."
	go build -o bin/app ./...

# Generate a new SQL migration
# Usage: make generate name=your_migration_name
generate:
	@echo "Generating SQL migration: $(name)"
	$(LIQUIBASE) \
		--changeLogFile=$(CHANGELOG_FILE) \
		generateChangeLog \
		--outputFile=$(MIGRATION_DIR)/$(shell date +"%Y%m%d%H%M%S")_$(name).sql

# Run migrations using Liquibase
migrate:
	@echo "Running Liquibase migrations..."
	$(LIQUIBASE) \
		--changeLogFile=$(CHANGELOG_FILE) \
		update

# Run docker-compose
up:
	docker-compose --env-file .env up -d

# Stop containers
down:
	docker-compose down

# Recreate database and run all migrations
reset:
	docker-compose down -v
	docker-compose --env-file .env up -d
	sleep 5
	make migrate
