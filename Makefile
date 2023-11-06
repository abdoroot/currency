run:
	go run main.go

proto:
	protoc --go_out=./currency --go-grpc_out=./currency  ./proto.proto