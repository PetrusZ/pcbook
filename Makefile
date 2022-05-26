.PHONY: gen clean server client test cert

gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb --grpc-gateway_out ./pb --openapiv2_out ./openapiv2

clean:
	rm pb/*.go

server1:
	go run cmd/server/main.go -port 50051

server2:
	go run cmd/server/main.go -port 50052

server1-tls:
	go run cmd/server/main.go -port 50051 -tls

server2-tls:
	go run cmd/server/main.go -port 50052 -tls

server:
	go run cmd/server/main.go -port 8080

rest:
	go run cmd/server/main.go -port 8081 -type rest -endpoint 127.0.0.1:8080

client:
	go run cmd/client/main.go -address 127.0.0.1:8080

client-tls:
	go run cmd/client/main.go -address 127.0.0.1:8080 -tls

test:
	go test -cover -race ./...

cert:
	cd cert; ./gen.sh; cd ..