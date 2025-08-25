FROM golang:alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum from the 'backend' subdirectory of the build context
COPY backend/go.mod backend/go.sum ./

RUN go mod download

# Copy the docs directory from the project root
COPY docs ./docs
# Copy the rest of the backend source code
COPY backend/ .

RUN go build -o main ./main_local.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs

EXPOSE 8080

CMD ["./main"]
