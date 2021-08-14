package test

import (
	"asanpardakht/profile/api/proto"
	"asanpardakht/profile/api/rpc_handler"
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	serverImpl := rpc_handler.NewProfileServer()
	proto.RegisterProfileServer(s, serverImpl)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestPurchasePermission_GoodGuy(t *testing.T) {
	type arg struct {
		cellPhone    string
		purchaseType proto.PurchasePermissionRequest_PurchaseType
	}

	tests := []struct {
		name string
		arg
		errorMessage string
	}{
		{
			name: "success",
			arg: arg{
				cellPhone:    "09387207022",
				purchaseType: proto.PurchasePermissionRequest_ByWallet,
			},
			errorMessage: "",
		},
		{
			name: "cell phone is required",
			arg: arg{
				cellPhone:    "",
				purchaseType: proto.PurchasePermissionRequest_ByWallet,
			},
			errorMessage: "rpc error: code = Unknown desc = user cell phone is required",
		},
		{
			name: "purchase type is required",
			arg: arg{
				cellPhone:    "09387207022",
				purchaseType: proto.PurchasePermissionRequest_None,
			},
			errorMessage: "rpc error: code = Unknown desc = purchase type is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
			if err != nil {
				t.Fatalf("Failed to dial bufnet: %v", err)
			}
			defer conn.Close()
			client := proto.NewProfileClient(conn)
			resp, err := client.PurchasePermission(ctx, &proto.PurchasePermissionRequest{
				UserCellPhone: tt.arg.cellPhone,
				PurchaseType:  tt.arg.purchaseType,
			})
			if tt.errorMessage == "" {
				assert.Nil(t, err)
				assert.NotNil(t, resp)
			} else {
				assert.Equal(t, tt.errorMessage, err.Error())
				assert.Nil(t, resp)
			}
		})
	}
}