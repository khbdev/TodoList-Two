FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .


RUN go build -o notes-service ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/notes-service .

EXPOSE 50052

CMD ["./notes-service"]