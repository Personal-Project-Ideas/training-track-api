# Use the official Go 1.24 image with Alpine as the base
FROM golang:1.24-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy everything from the project root
COPY . .

# Download Go module dependencies
RUN go mod download

# Build the Go application from the main package
RUN go build -o main ./api/cmd/server

# Set the default command to run the application binary
CMD ["./main"]
