# Build stage
FROM golang:1.24-alpine AS builder

# Install swag for generating Swagger docs
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Set the working directory
WORKDIR /app

# Copy the Go module files first (for better caching)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Generate Swagger documentation
RUN swag init -g main.go  

# Ensure the docs folder is copied
COPY docs /app/docs

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o backend .

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/backend .

# Copy the docs folder (Swagger JSON)
COPY --from=builder /app/docs /app/docs

# Expose port 8080 (matching docker-compose)
EXPOSE 8080

# Run the Go app
CMD ["./backend"]
