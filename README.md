# Notes-Service
Key-Value Notes Service using Golang, Protobuf, gRPC

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
