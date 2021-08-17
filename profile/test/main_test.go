package test

import (
	"asanpardakht/profile/api/proto"
	"asanpardakht/profile/api/rpc_handler"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"os"
	"testing"
)

var (
	client   proto.ProfileClient
	closeFns func()
)

func TestMain(m *testing.M) {
	client, closeFns = runFakeProfileService()
	code := m.Run()
	closeFns()
	os.Exit(code)
}

func runFakeProfileService() (proto.ProfileClient, func()) {
	serverImpl := rpc_handler.NewProfileServer()
	server := grpc.NewServer()
	proto.RegisterProfileServer(server, serverImpl)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", FakeProfileServicePort))
	if err != nil {
		panic(err)
	}
	go func() {
		if err := server.Serve(lis); err != nil {
			panic(err)
		}
	}()

	conn, err := grpc.Dial(fmt.Sprintf(":%d", FakeProfileServicePort), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	var closeFn = func() {
		conn.Close()
		server.Stop()
	}

	return proto.NewProfileClient(conn), closeFn
}
