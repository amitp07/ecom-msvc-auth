BINARY_NAME=auth-msvc

build:
	@echo "Building application..."
	go build -o $(BINARY_NAME) ./cmd/app

run: build
	@echo "Running the app..."
	./${BINARY_NAME}
