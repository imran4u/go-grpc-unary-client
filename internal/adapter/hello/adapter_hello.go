package hello

import (
	"context"

	"github.com/imran4u/go-grpc-proto/protogen/go/hello"
	"github.com/imran4u/go-grpc-unary-client/internal/port"
	"google.golang.org/grpc"
)

type HelloAdapter struct {
	helloClient port.HelloClientPort
}

func NewHelloAdapter(conn *grpc.ClientConn) (*HelloAdapter, error) {
	client := hello.NewHelloServiceClient(conn)
	return &HelloAdapter{
		helloClient: client,
	}, nil
}

func (a *HelloAdapter) SayHello(name string) string {
	req := &hello.HelloRequest{
		Name: name,
	}
	res, err := a.helloClient.SayHello(context.Background(), req)
	if err != nil {
		return err.Error()
	}
	return res.Greet
}
