FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY ../../api/go.mod ./
COPY ../../api/go.sum ./
RUN go mod download

COPY ../../api ./
RUN go build -o main .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
