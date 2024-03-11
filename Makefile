protoc:
	protoc --proto_path=proto proto/*.proto --go_out=$(dst) --go-grpc_out=$(dst)

run:
	@echo "run server..."
	go run cmd/main.go