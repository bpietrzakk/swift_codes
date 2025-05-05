# Use an official Go image as the base image, specifying the platform for ARM compatibility
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY .env .
COPY internal/data/ /app/internal/data/

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd

# Use a minimal base image for the final container, specifying the platform
FROM debian:bullseye-slim

# Install certificates if your app makes HTTPS requests
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Set the working directory in the container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env . 
COPY --from=builder /app/internal/data/ /internal/data/


EXPOSE 8080

CMD ["./main"]