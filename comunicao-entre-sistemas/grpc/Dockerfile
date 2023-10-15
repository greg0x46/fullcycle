FROM golang:latest

WORKDIR /go/src/app
COPY . .

RUN apt-get update \
    && apt install -y protobuf-compiler \
    && go get google.golang.org/protobuf/cmd/protoc-gen-go \
    && go get google.golang.org/grpc/cmd/protoc-gen-go-grpc \
    && go get google.golang.org/grpc


EXPOSE 50051:50051