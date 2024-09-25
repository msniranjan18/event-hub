# Define variables
APP_NAME := event-hub-backend
DOCKER_IMAGE := event-hub-backend
DOCKER_TAG := 1.0.0
BUILD_DIR := ${PWD}/build
CONTAINER_NAME := event-hub-backend
# Default target
all: build

# Build the Go application
build:
	# Note: This build will only run on this machine, for docker build will use multi stage Dockerfile to build a machine comaptible binary.
	@echo "Building the Go application..." 
	mkdir -p $(BUILD_DIR)
	# GOOS=linux GOARCH=amd64 go build -a -o $(BUILD_DIR)/$(APP_NAME) main.go
	go mod tidy
	go build -o $(BUILD_DIR)/$(APP_NAME) main.go

# Build Docker image
docker-build:
	@echo "Building Docker image..."
	docker build -t \
	--build-arg GOOS=linux \
	--build-arg GOARCH=amd64 \
	--build-arg CGO_ENABLED=1 \
	--build-arg APP_ENV=staging \
	--build-arg PORT=8080 \
	--build-arg APP_NAME=event-hub-backend \
	--build-arg BUILD_VERSION=1.0.0 \
	$(DOCKER_IMAGE):$(DOCKER_TAG) -f docker/Dockerfile .

# Run Docker container
docker-run:
	@echo "Stopping and removing existing Docker container..."
	# Stop and remove any existing container with the same name
	CONTAINER_ID=$$(docker ps -aq -f name=$(CONTAINER_NAME)); \
	if [ -n "$$CONTAINER_ID" ]; then \
		docker stop $$CONTAINER_ID || true; \
		docker rm $$CONTAINER_ID || true; \
	fi
	@echo "Running Docker container..."
	# running in detach mode
	docker run -d -p 8080:8080 --name $(APP_NAME) $(DOCKER_IMAGE):$(DOCKER_TAG)

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)/$(APP_NAME)
	CONTAINER_ID=$$(docker ps -aq -f name=$(CONTAINER_NAME)); \
	if [ -n "$$CONTAINER_ID" ]; then \
		docker stop $$CONTAINER_ID || true; \
		docker rm -f $$CONTAINER_ID || true; \
	fi
	docker rmi $(DOCKER_IMAGE):$(DOCKER_TAG)

# Run all tests
test:
	@echo "Running tests..."
	go test ./...

# Apply Kubernetes configurations
k8s-deploy:
	@echo "Applying Kubernetes configurations..."
	kubectl apply -f kubernetes/namespace.yaml
	kubectl apply -f kubernetes/

# Help message
help:
	@echo "Makefile for managing the backend application"
	@echo ""
	@echo "Usage:"
	@echo "  make build          Build the Go application"
	@echo "  make docker-build   Build the Docker image"
	@echo "  make docker-run     Run the Docker container"
	@echo "  make clean          Clean build artifacts and Docker image"
	@echo "  make test           Run tests"
	@echo "  make k8s-deploy     Apply Kubernetes configurations"
	@echo "  make help           Show this help message"
