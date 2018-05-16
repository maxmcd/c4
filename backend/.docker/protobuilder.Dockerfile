# Dockerfile for compiling protos
FROM golang:1.10

# Install unzip
RUN apt-get update && apt-get install -y unzip

# Install protoc version 3.3.0
RUN wget https://github.com/google/protobuf/releases/download/v3.3.0/protoc-3.3.0-linux-x86_64.zip
RUN unzip protoc-3.3.0-linux-x86_64.zip

# Install protoc-gen-go
RUN go get -u github.com/golang/protobuf/protoc-gen-go
# Install protoc-gen-swagger
RUN go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
# Install protoc-gen-grpc-gateway
RUN go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

# Install go-bindata
RUN go get -u github.com/jteeuwen/go-bindata/...
