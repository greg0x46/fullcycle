FROM golang:latest

WORKDIR /go/src/app
COPY . .

RUN apt-get update \
    && apt install -y protobuf-compiler 

RUN  go get google.golang.org/protobuf/cmd/protoc-gen-go \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 \
    && go get github.com/mattn/go-sqlite3 \
    && go install github.com/ktr0731/evans@latest \
    && go mod tidy

RUN export PATH="$PATH:$(go env GOPATH)/bin"

EXPOSE 50051:50051