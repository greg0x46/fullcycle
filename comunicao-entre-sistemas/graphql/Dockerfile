FROM golang:latest

WORKDIR /app
COPY . .

RUN go mod tidy

EXPOSE 9090

# CMD go run cmd/server/server.go