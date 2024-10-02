# Build stage
FROM golang:1.22-bullseye AS builder

# Build-time argument (can be passed during build with --build-arg)
ARG APP_NAME=eventhub-backend
ARG APP_ENV=staging
ARG BUILD_VERSION=2.0.0
ARG CGO_ENABLED=1
ARG GOOS=linux
ARG GOARCH=amd64
ARG PORT=8080

# Install build dependencies
RUN apt-get update && apt-get install -y build-essential

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
# Seperate layer for dependancies to reduce the build time.
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=${CGO_ENABLED} GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${APP_NAME} main.go

# Final stage
FROM debian:bullseye-slim

# Set runtime environment variables
ENV APP_NAME=${APP_NAME}
ENV APP_ENV=${APP_ENV}
ENV BUILD_VERSION=${BUILD_VERSION}
ENV PORT=${PORT}

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
#COPY --from=builder /app/${APP_NAME} /bin/${APP_NAME}
COPY --from=builder /app/eventhub-backend /bin/eventhub-backend

# make it executable
RUN chmod +x /bin/eventhub-backend

# Expose port 8080 to the outside world
EXPOSE ${PORT}

# Command to run the executable
CMD ["/bin/eventhub-backend"]
#CMD ./bin/${APP_NAME}
