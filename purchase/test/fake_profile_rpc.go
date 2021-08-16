package test

import (
	"asanpardakht/profile/api/proto"
	"context"
	"math/rand"
	"time"
)

type FakeProfileService struct{}

func (p *FakeProfileService) PurchasePermission(_ context.Context, request *proto.PurchasePermissionRequest) (*proto.PurchasePermissionResponse, error) {
	if request.GetUserCellPhone() == "" {
		return &proto.PurchasePermissionResponse{
			Error:      proto.PurchasePermissionResponse_UserPhoneValidation,
			Permission: false,
		}, nil
	}

	if request.GetPurchaseType() == proto.PurchasePermissionRequest_None {
		return &proto.PurchasePermissionResponse{
			Error:      proto.PurchasePermissionResponse_PurchaseTypeValidation,
			Permission: false,
		}, nil
	}

	source := rand.NewSource(time.Now().UnixNano())
	customRand := rand.New(source)
	randInt := customRand.Int31()

	per := randInt%2 == 0 || randInt%3 == 0 || randInt%5 == 0
	return &proto.PurchasePermissionResponse{
		Permission: per,
	}, nil
}
