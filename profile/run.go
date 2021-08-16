package profile

import (
	"asanpardakht/profile/api/proto"
	"asanpardakht/profile/api/rpc_handler"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

const (
	profileRpcPort = 70051
)

func RunProfileService() (func(), int, error) {
	// Initialize the generated Profile service.
	svc := rpc_handler.NewProfileServer()

	// Create a new gRPC server.
	server := grpc.NewServer()

	// Register the User service with the server.
	proto.RegisterProfileServer(server, svc)

	// Open port profileRpcPort for listening to traffic.
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", profileRpcPort))
	if err != nil {
		return nil, 0, err
	}

	// Listen for traffic indefinitely.
	if err := server.Serve(lis); err != nil {
		return nil, 0, err
	}

	return server.Stop, profileRpcPort, nil
}
