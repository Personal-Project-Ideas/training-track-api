# Use the official Go 1.24 image with Alpine as the base
FROM golang:1.24-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy Go module files first to leverage Docker layer caching
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application from the main package
RUN go build -o main ./cmd/server

# Set the default command to run the application binary
CMD ["./main"]
