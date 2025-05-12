# Notes-Service

A Key-Value Notes Service built using **Golang**, **Protobuf**, and **gRPC**.

## Table of Contents

- [Architecture Overview](#architecture-overview)
- [Technology Stack](#technology-stack)
- [Setup Instructions](#setup-instructions)
- [Protoc Compilation](#protoc-compilation)
- [API Usage](#api-usage)
- [License](#license)

---

## Architecture Overview


[ REST Client (Postman/cURL) ] 
|
V
[ gRPC-Gateway (runtime.NewServeMux) (port 8080)]
|
V
[ gRPC Server (port 50051) ]
|
V
[ etcd (key-value store) ]

## Technology Stack

- **Language**: Go (Golang)
- **Protocol**: gRPC
- **IDL**: Protocol Buffers (Protobuf)
- **Storage**: etcd (Key-Value Store)
- **Gateway**: gRPC-Gateway for HTTP/JSON -> gRPC translation

## Protoc Compilation

Make sure you compile `.proto` files using the correct include paths:

```bash
protoc -I protos -I third_party \
  --go_out=. \
  --go-grpc_out=. \
  --grpc-gateway_out=. \
  --openapiv2_out=. \
  protos/notes.proto
