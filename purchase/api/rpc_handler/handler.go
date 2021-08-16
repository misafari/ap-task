package rpc_handler

import (
	profileProto "asanpardakht/profile/api/proto"
	"asanpardakht/purchase/api/proto"
	"context"
)

type PurchaseService struct {
	profileClient profileProto.ProfileClient
}

func NewPurchaseService(profileClient profileProto.ProfileClient) *PurchaseService {
	return &PurchaseService{
		profileClient: profileClient,
	}
}

func (p *PurchaseService) PayByCart(ctx context.Context, in *proto.PayByCartRequest) (*proto.PurchaseResponse, error) {
	profileResponse, err := p.profileClient.PurchasePermission(ctx, &profileProto.PurchasePermissionRequest{
		UserCellPhone: "",
		PurchaseType:  0,
	})
	if err != nil {
		// log err
		return &proto.PurchaseResponse{
			UserPhoneNo: in.GetUserPhoneNo(),
			Error:       proto.PurchaseResponse_Unknown,
		}, nil
	}

	if !profileResponse.Permission {
		return &proto.PurchaseResponse{
			UserPhoneNo: in.GetUserPhoneNo(),
			Error:       proto.PurchaseResponse_Permission,
		}, nil
	}

	if !PayByCartIncomeValidation(in) {
		return &proto.PurchaseResponse{
			UserPhoneNo: in.GetUserPhoneNo(),
			Error:       proto.PurchaseResponse_Info,
		}, nil
	}

	resp := &proto.PurchaseResponse{
		UserPhoneNo: in.GetUserPhoneNo(),
		Error:       MakeFakeDecision(),
	}

	return resp, nil
}
