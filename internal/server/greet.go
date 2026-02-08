package server

import (
	"context"
	"fmt"

	connect "connectrpc.com/connect"
	greetv1 "github.com/dxh9845/explore-ast-grep/gen/greet/v1"
)

type GreetServer struct{}

func (s *GreetServer) Greet(
	ctx context.Context,
	req *connect.Request[greetv1.GreetRequest],
) (*connect.Response[greetv1.GreetResponse], error) {
	name := req.Msg.Name
	if name == "" {
		name = "World"
	}
	return connect.NewResponse(&greetv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", name),
	}), nil
}
