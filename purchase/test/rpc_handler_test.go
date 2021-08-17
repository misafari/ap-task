package test

import (
	"asanpardakht/purchase/api/proto"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPayByCart_GoodGuy(t *testing.T) {
	type arg struct {
		userCellPhoneNumber string
		cartNumber          string
		amount              int32
	}

	tests := []struct {
		name string
		arg
		errorMessage proto.PurchaseResponse_Error
	}{
		{
			name: "Permission Denied",
			arg: arg{
				// 09121231234 allowed user (fake)
				userCellPhoneNumber: "09121231235",
				cartNumber:          "",
				amount:              0,
			},
			errorMessage: proto.PurchaseResponse_Permission,
		},
		{
			name: "Validation Err",
			arg: arg{
				userCellPhoneNumber: "09121231234",
				cartNumber:          "",
				amount:              0,
			},
			errorMessage: proto.PurchaseResponse_Info,
		},
		{
			name: "Validation Err",
			arg: arg{
				userCellPhoneNumber: "09121231234",
				cartNumber:          "1234123412341234",
				amount:              0,
			},
			errorMessage: proto.PurchaseResponse_Info,
		},
		{
			name: "Random Error",
			arg: arg{
				userCellPhoneNumber: "09121231234",
				cartNumber:          "1234123412341234",
				amount:              10000,
			},
			errorMessage: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := client.PayByCart(context.Background(), &proto.PayByCartRequest{
				UserPhoneNo: tt.userCellPhoneNumber,
				CartNumber:  tt.cartNumber,
				Amount:      tt.amount,
			})

			assert.Nil(t, err)
			assert.NotNil(t, resp)
			if tt.errorMessage != 0 {
				assert.Equal(t, resp.Error, tt.errorMessage)
			}
		})
	}
}

func TestPayByWallet_GoodGuy(t *testing.T) {
	type arg struct {
		userCellPhoneNumber string
		walletNumber        string
		amount              int32
	}

	tests := []struct {
		name string
		arg
		errorMessage proto.PurchaseResponse_Error
	}{
		{
			name: "Permission Denied",
			arg: arg{
				// 09121231234 allowed user (fake)
				userCellPhoneNumber: "09121231235",
				walletNumber:        "123",
				amount:              0,
			},
			errorMessage: proto.PurchaseResponse_Permission,
		},
		{
			name: "Validation Err",
			arg: arg{
				userCellPhoneNumber: "09121231234",
				walletNumber:        "",
				amount:              0,
			},
			errorMessage: proto.PurchaseResponse_Info,
		},
		{
			name: "Validation Err",
			arg: arg{
				userCellPhoneNumber: "09121231234",
				walletNumber:        "123",
				amount:              0,
			},
			errorMessage: proto.PurchaseResponse_Info,
		},
		{
			name: "Random Error",
			arg: arg{
				userCellPhoneNumber: "09121231234",
				walletNumber:        "1234123412341234",
				amount:              10000,
			},
			errorMessage: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := client.PayByWallet(context.Background(), &proto.PayByWalletRequest{
				UserPhoneNo:  tt.userCellPhoneNumber,
				WalletNumber: tt.walletNumber,
				Amount:       tt.amount,
			})

			assert.Nil(t, err)
			assert.NotNil(t, resp)
			if tt.errorMessage != 0 {
				assert.Equal(t, resp.Error, tt.errorMessage)
			}
		})
	}
}

func TestPayByCredit_GoodGuy(t *testing.T) {
	type arg struct {
		userCellPhoneNumber string
		creditId            string
		amount              int32
	}

	tests := []struct {
		name string
		arg
		errorMessage proto.PurchaseResponse_Error
	}{
		{
			name: "Permission Denied",
			arg: arg{
				// 09121231234 allowed user (fake)
				userCellPhoneNumber: "09121231235",
				creditId:            "123",
				amount:              0,
			},
			errorMessage: proto.PurchaseResponse_Permission,
		},
		{
			name: "Validation Err",
			arg: arg{
				userCellPhoneNumber: "09121231234",
				creditId:            "",
				amount:              0,
			},
			errorMessage: proto.PurchaseResponse_Info,
		},
		{
			name: "Validation Err",
			arg: arg{
				userCellPhoneNumber: "09121231234",
				creditId:            "123",
				amount:              0,
			},
			errorMessage: proto.PurchaseResponse_Info,
		},
		{
			name: "Random Error",
			arg: arg{
				userCellPhoneNumber: "09121231234",
				creditId:            "1234123412341234",
				amount:              10000,
			},
			errorMessage: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := client.PayByCredit(context.Background(), &proto.PayByCreditRequest{
				UserPhoneNo: tt.userCellPhoneNumber,
				CreditId:    tt.creditId,
				Amount:      tt.amount,
			})

			assert.Nil(t, err)
			assert.NotNil(t, resp)
			if tt.errorMessage != 0 {
				assert.Equal(t, resp.Error, tt.errorMessage)
			}
		})
	}
}
