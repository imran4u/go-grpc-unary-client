package port

import (
	"github.com/imran4u/go-grpc-proto/protogen/go/hello"
)

type HelloClientPort interface {
	Hello(request *hello.HelloRequest) (*hello.HelloResponse, error)
}
