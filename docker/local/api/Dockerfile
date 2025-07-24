# Use official Go Alpine image as the build environment
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY ./api/go.mod ./api/go.sum ./

RUN go mod download

# Copy only the local code folder containing main.go for local dev
COPY ./api/local/ ./local/

WORKDIR /app/local

# Build the local version of the API binary
RUN go build -o /api-local .

# Final minimal image with just the compiled binary
FROM alpine:latest

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /api-local /app/api-local

EXPOSE 8080

# Run the binary by default
CMD ["/app/api-local"]
