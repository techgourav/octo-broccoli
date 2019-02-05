package server

import (
	"net"

	todopb "octo-broccoli/protocol/go"

	"google.golang.org/grpc"
)

type server struct{}

// Start : Start a grpc server
func Start() error {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	todopb.RegisterTodoServiceServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		return err
	}

	return nil
}
