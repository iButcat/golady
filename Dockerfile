FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum to leverage Docker caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -o /app/github-notify ./cmd/main.go

# Use a smaller image for the final container
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata sqlite

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/github-notify .

# Create a volume for persistent data
VOLUME /app/data

# Set environment variables with defaults
ENV DB_PATH=/app/data/github-notify.db
ENV SERVER_PORT=8080

# Expose the server port
EXPOSE 8080

# Run the application
CMD ["/app/github-notify"]