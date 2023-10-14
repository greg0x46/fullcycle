FROM golang:latest

WORKDIR /usr/src/app
COPY . .

RUN set -eux; \
	export GOROOT="$(go env GOROOT)"; \
	./run-tool-which-requires-GOROOT-set.sh

RUN go mod tidy

EXPOSE 9090

CMD go run cmd/server/server.go