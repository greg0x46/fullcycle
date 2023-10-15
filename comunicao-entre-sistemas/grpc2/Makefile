current_dir := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

build:
	docker build -t grpc2 .
up:
	docker run -d --name grpc2 -v "$(current_dir)/.:/go/src/app" grpc2
run-bash: 
	docker run -it --name grpc2 -v "$(current_dir)/.:/go/src/app" grpc2 bash
bash:
	docker exec -it grpc2 bash
proto-gen:
	docker exec grpc2 protoc --go_out=. --go-grpc_out=. proto/course_category.proto