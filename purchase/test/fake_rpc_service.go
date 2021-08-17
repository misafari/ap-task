package test

import (
	profileProto "asanpardakht/profile/api/proto"
	"context"
)

type FakeProfileService struct{}

func (p *FakeProfileService) PurchasePermission(_ context.Context, request *profileProto.PurchasePermissionRequest) (*profileProto.PurchasePermissionResponse, error) {
	return &profileProto.PurchasePermissionResponse{
		Permission: request.GetUserCellPhone() == "09121231234",
	}, nil
}
