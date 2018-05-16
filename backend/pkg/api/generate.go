//go:generate protoc --plugin=protoc-gen-go=$GOPATH/bin/protoc-gen-go --go_out=plugins=grpc:./ --plugin=protoc-gen-grpc-gateway=$GOPATH/bin/protoc-gen-grpc-gateway --swagger_out=logtostderr=true:./ --grpc-gateway_out=logtostderr=true:./ -I./ -I../../third_party/googleapis auth_service.proto

package api
