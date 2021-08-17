package test

import (
	profileProto "asanpardakht/profile/api/proto"
	purchaseProto "asanpardakht/purchase/api/proto"
	"asanpardakht/purchase/api/rpc_handler"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"os"
	"testing"
	"time"
)

const (
	FakeProfileServicePort  = 60061
	FakePurchaseServicePort = 60063
)

var (
	client   purchaseProto.PurchaseClient
	closeFns func()
)

func TestMain(m *testing.M) {
	client, closeFns = runPurchaseRpc()
	code := m.Run()
	closeFns()
	os.Exit(code)
}

func runFakeProfileService() (profileProto.ProfileClient, func()) {
	serverImpl := &FakeProfileService{}
	server := grpc.NewServer()
	profileProto.RegisterProfileServer(server, serverImpl)
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

	return profileProto.NewProfileClient(conn), closeFn
}

func runPurchaseRpc() (purchaseProto.PurchaseClient, func()) {
	profileRpcClient, profileRpcClientCloseFn := runFakeProfileService()
	serverImpl := rpc_handler.NewPurchaseService(profileRpcClient, 2*time.Second)
	server := grpc.NewServer()
	purchaseProto.RegisterPurchaseServer(server, serverImpl)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", FakePurchaseServicePort))
	if err != nil {
		panic(err)
	}

	go func() {
		if err := server.Serve(lis); err != nil {
			panic(err)
		}
	}()

	conn, err := grpc.Dial(fmt.Sprintf(":%d", FakePurchaseServicePort), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	var closeFn = func() {
		conn.Close()
		server.Stop()
		profileRpcClientCloseFn()
	}

	return purchaseProto.NewPurchaseClient(conn), closeFn
}
