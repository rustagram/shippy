build:
	protoc --micro_out=. --go_out=. proto/consignment/consignment.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t shippy .
run:
	docker run -p 50051:50051 shippy
	    -e MICRO_SERVER_ADDRESS=:50051 \
	    shippy