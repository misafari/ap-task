package test

import (
	"asanpardakht/profile/api/proto"
	"asanpardakht/profile/api/rpc_handler"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"net"
	"testing"
)

const (
	FakeProfileServicePort = 60062
)

func TestPurchasePermission_GoodGuy(t *testing.T) {
	type arg struct {
		cellPhone    string
		purchaseType proto.PurchasePermissionRequest_PurchaseType
	}

	tests := []struct {
		name string
		arg
		errorMessage proto.PurchasePermissionResponse_PurchasePermissionResponseError
	}{
		{
			name: "success",
			arg: arg{
				cellPhone:    "09387207022",
				purchaseType: proto.PurchasePermissionRequest_ByWallet,
			},
			errorMessage: proto.PurchasePermissionResponse_None,
		},
		{
			name: "cell phone is required",
			arg: arg{
				cellPhone:    "",
				purchaseType: proto.PurchasePermissionRequest_ByWallet,
			},
			errorMessage: proto.PurchasePermissionResponse_UserPhoneValidation,
		},
		{
			name: "purchase type is required",
			arg: arg{
				cellPhone:    "09387207022",
				purchaseType: proto.PurchasePermissionRequest_None,
			},
			errorMessage: proto.PurchasePermissionResponse_PurchaseTypeValidation,
		},
	}

	client, closeFns := runFakeProfileService()
	defer closeFns()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := client.PurchasePermission(context.Background(), &proto.PurchasePermissionRequest{
				UserCellPhone: tt.arg.cellPhone,
				PurchaseType:  tt.arg.purchaseType,
			})
			assert.Nil(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, tt.errorMessage, resp.Error)
		})
	}
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
