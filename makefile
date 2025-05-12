PROTO_SRC = protos/notes.proto
GEN_GO_DIR = ./gen/go
GEN_SWAGGER_DIR = ./gen/swagger

proto:
	mkdir -p $(GEN_GO_DIR) $(GEN_SWAGGER_DIR)
	protoc \
		-I protos \
		-I third-party \
		--go_out=$(GEN_GO_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(GEN_GO_DIR) \
		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=$(GEN_GO_DIR) \
		--grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=$(GEN_SWAGGER_DIR) \
		$(PROTO_SRC)
