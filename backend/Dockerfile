FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Change the build target to main_local.go
RUN go build -o main ./main_local.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

# Execute the new binary
CMD ["./main"]
