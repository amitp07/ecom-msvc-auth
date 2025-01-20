BINARY_NAME=auth-msvc

build:
	@echo "Building application..."
	go build -o $(BINARY_NAME) ./cmd/api

run: build
	@echo "Running the app..."
	./${BINARY_NAME}

MIGRATE_CMD=migrate -path ./cmd/migrations -database $(PG_URL)

migrate-up:
	$(MIGRATE_CMD) up


