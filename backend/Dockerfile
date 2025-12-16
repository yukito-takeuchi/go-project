# Development stage with Air for hot reload
FROM golang:1.21-alpine AS development

WORKDIR /app

# Install Air for hot reload
RUN go install github.com/cosmtrek/air@latest

# Copy go mod files
COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Expose port
EXPOSE 8080

# Air will be run via docker-compose command
CMD ["air", "-c", ".air.toml"]

# Production stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final production stage
FROM alpine:latest AS production

WORKDIR /root/

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Copy binary from builder
COPY --from=builder /app/main .

# Expose port
EXPOSE 8080

# Run the application
CMD ["./main"]
