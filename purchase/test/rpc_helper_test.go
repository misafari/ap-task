package test

import (
	"asanpardakht/purchase/api/proto"
	"asanpardakht/purchase/api/rpc_handler"
	"testing"
)

func Test_cartNumberValidation(t *testing.T) {
	tests := []struct {
		name           string
		cardNumber     string
		expectedResult bool
	}{
		{
			name:           "Should Be Ok",
			cardNumber:     "1234123412341234",
			expectedResult: true,
		},
		{
			name:           "Should Be Failed",
			cardNumber:     "123412341234123",
			expectedResult: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rpc_handler.CartNumberValidation(tt.cardNumber); got != tt.expectedResult {
				t.Errorf("cartNumberValidation() = %v, want %v", got, tt.expectedResult)
			}
		})
	}
}

func Test_payByCartIncomeValidation(t *testing.T) {
	type args struct {
		in *proto.PayByCartRequest
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should Be Ok",
			args: args{
				in: &proto.PayByCartRequest{
					UserPhoneNo: "09121231234",
					CartNumber:  "1234123412341234",
					Amount:      10000,
				},
			},
			want: true,
		},
		{
			name: "Phone is required",
			args: args{
				in: &proto.PayByCartRequest{
					UserPhoneNo: "",
					CartNumber:  "1234123412341234",
					Amount:      10000,
				},
			},
			want: false,
		},
		{
			name: "Cart is required",
			args: args{
				in: &proto.PayByCartRequest{
					UserPhoneNo: "09121231234",
					CartNumber:  "",
					Amount:      10000,
				},
			},
			want: false,
		},
		{
			name: "Cart is not valid",
			args: args{
				in: &proto.PayByCartRequest{
					UserPhoneNo: "09121231234",
					CartNumber:  "12341234123412",
					Amount:      10000,
				},
			},
			want: false,
		},
		{
			name: "Amount is required",
			args: args{
				in: &proto.PayByCartRequest{
					UserPhoneNo: "09121231234",
					CartNumber:  "1234123412341234",
					Amount:      0,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rpc_handler.PayByCartIncomeValidation(tt.args.in); got != tt.want {
				t.Errorf("payByCartIncomeValidation() = %v, want %v", got, tt.want)
			}
		})
	}
}
