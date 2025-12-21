.PHONY: test test-unit test-bdd

NOW := $(shell date '+%Y%m%d%H%M%S')

# Configurable variables (can be set in environment or .env)
DATABASE_HOST ?= localhost
DATABASE_PORT ?= 5432
DATABASE_NAME ?= trainingLog
DATABASE_USER ?= admin
DATABASE_PASSWD ?= s9X@7vP2wQz!F4mL

# ------------------------
# API Build and Run Commands
# ------------------------

# Build the API image for local development
api-build-local:
	docker build -f docker/local/api/Dockerfile -t local-api ./api

# Run the API container locally with DB env vars
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

# Run the API locally using Go (development mode)
api-dev:
	go run ./api/local

# Build the production-ready API image
api-build-prod:
	docker build -f docker/Dockerfile -t prod-api .

# Run the production API container
api-run-prod:
	docker run --rm -p 8080:8080 prod-api

# Run only unit tests
test-unit:
	go test ./test/unit/... -count=1 -v


#Run only integration tests
test-bdd:
	go test ./test/bdd/... -count=1

# Run all tests
test:
	go test ./test/... -count=1 -v
