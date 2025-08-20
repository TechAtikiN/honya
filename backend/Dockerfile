FROM golang:1.24-alpine AS builder

# Set the working directory to /src
WORKDIR /app

# Copy the Go module files first (for better caching)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the db directory containing the Go files
COPY . .

# Build the Go app 
RUN go build -o backend

# Expose port 8000
EXPOSE 8000

# Run the Go app
CMD ["./backend"]