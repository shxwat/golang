FROM golang:1.26.8-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

# Reach PostgreSQL running on the Docker Desktop host.
ENV DB_HOST=host.docker.internal \
    DB_PORT=5432

EXPOSE 8080

CMD ["./main"]
