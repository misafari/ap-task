package test

import (
	"asanpardakht/purchase/api/proto"
	"asanpardakht/purchase/api/rpc_handler"
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

	serverImpl := rpc_handler.NewPurchaseService(nil)
	proto.RegisterPurchaseServer(s, serverImpl)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestPayByCart_GoodGuy(t *testing.T) {
	type arg struct {
		userCellPhoneNumber string
		cartNumber          string
		amount              int32
	}

	tests := []struct {
		name string
		arg
		errorMessage string
	}{
		{
			name: "Should Be Ok",
			arg: arg{
				userCellPhoneNumber: "",
				cartNumber:          "",
				amount:              0,
			},
			errorMessage: "",
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
			client := proto.NewPurchaseClient(conn)
			resp, err := client.PayByCart(ctx, &proto.PayByCartRequest{
				UserPhoneNo: tt.userCellPhoneNumber,
				CartNumber:  tt.cartNumber,
				Amount:      tt.amount,
			})

			assert.Nil(t, err)
			assert.NotNil(t, resp)
		})
	}
}
