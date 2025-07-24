NOW := $(shell date '+%Y%m%d%H%M%S')

# Configurable variables (can be set in environment or .env)
DATABASE_HOST ?= localhost
DATABASE_PORT ?= 5432
DATABASE_NAME ?= trainlog
DATABASE_USER ?= admin
DATABASE_PASSWD ?= admin123

LIQUIBASE_IMAGE := liquibase/liquibase

# Paths for migrations and changelog in the project root
MIGRATIONS_DIR := ./migrations/sql
CHANGELOG_FILE := ./migrations/changelog.yml

# ------------------------
# Liquibase Commands via Docker
# ------------------------

liquibase-diff:
	docker run --rm $(LIQUIBASE_IMAGE) \
		diff-changelog \
		--url="jdbc:postgresql://$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_NAME)" \
		--username=$(DATABASE_USER) \
		--password=$(DATABASE_PASSWD) \
		--referenceUrl="jdbc:postgresql://$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_NAME)" \
		--referenceUsername=$(DATABASE_USER) \
		--referencePassword=$(DATABASE_PASSWD)

liquibase-migrate:
	docker run --rm -v $(PWD)/migrations:/liquibase $(LIQUIBASE_IMAGE) \
		update \
		--url="jdbc:postgresql://$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_NAME)" \
		--username=$(DATABASE_USER) \
		--password=$(DATABASE_PASSWD) \
		--changeLogFile=/liquibase/changelog.yml

liquibase-rollback:
	docker run --rm -v $(PWD)/migrations:/liquibase $(LIQUIBASE_IMAGE) \
		rollback-count 1 \
		--url="jdbc:postgresql://$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_NAME)" \
		--username=$(DATABASE_USER) \
		--password=$(DATABASE_PASSWD) \
		--changeLogFile=/liquibase/changelog.yml

liquibase-update-sql:
	docker run --rm -v $(PWD)/migrations:/liquibase $(LIQUIBASE_IMAGE) \
		updateSQL \
		--url="jdbc:postgresql://$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_NAME)" \
		--username=$(DATABASE_USER) \
		--password=$(DATABASE_PASSWD) \
		--changeLogFile=/liquibase/changelog.yml \
		--outputFile=/liquibase/script.sql

liquibase-create-file:
	@touch ./migrations/sql/$(NOW)_$(name).sql

# ------------------------
# API Build and Run Commands
# ------------------------

api-build-local:
	docker build -f docker/local/api/Dockerfile -t local-api ./api

# Run API container locally, passing DB env vars via -e
api-run-local:
	docker run --rm \
		-e DATABASE_HOST=$(DATABASE_HOST) \
		-e DATABASE_PORT=$(DATABASE_PORT) \
		-e DATABASE_NAME=$(DATABASE_NAME) \
		-e DATABASE_USER=$(DATABASE_USER) \
		-e DATABASE_PASSWD=$(DATABASE_PASSWD) \
		-p 8080:8080 \
		--network trainlog-net \
		local-api


# Run API directly in Go (dev mode)
api-dev:
	go run ./api/local

api-build-prod:
	docker build -f docker/Dockerfile -t prod-api .

api-run-prod:
	docker run --rm -p 8080:8080 prod-api
