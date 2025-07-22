# Start from the official Go image for building
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .

# Download dependencies
RUN go mod download

# Build the Go app
RUN go build -o app .

# Start a minimal image for running
FROM alpine:latest
WORKDIR /app

# Install CA certificates (for HTTPS requests)
RUN apk --no-cache add ca-certificates

# Copy the built Go binary
COPY --from=builder /app/app .

# Copy .env file if present (for local dev, optional in prod)
COPY .env* ./

# Expose the default port (can be overridden by env)
EXPOSE 8080

# Set environment variables (can be overridden at runtime)
ENV PORT=8080
ENV DB_PATH=studentpm.db

# Run the app
CMD ["./app"] 