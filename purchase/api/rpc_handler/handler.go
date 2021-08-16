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
	if permissionDenied := p.checkPermission(ctx, in.GetUserPhoneNo(), profileProto.PurchasePermissionRequest_ByCart); permissionDenied != nil {
		return permissionDenied, nil
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

func (p *PurchaseService) PayByWallet(ctx context.Context, in *proto.PayByWalletRequest) (*proto.PurchaseResponse, error) {
	if permissionDenied := p.checkPermission(ctx, in.GetUserPhoneNo(), profileProto.PurchasePermissionRequest_ByWallet); permissionDenied != nil {
		return permissionDenied, nil
	}

	if !PayByWalletIncomeValidation(in) {
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

func (p *PurchaseService) PayByCredit(ctx context.Context, in *proto.PayByCreditRequest) (*proto.PurchaseResponse, error) {
	if permissionDenied := p.checkPermission(ctx, in.GetUserPhoneNo(), profileProto.PurchasePermissionRequest_ByWallet); permissionDenied != nil {
		return permissionDenied, nil
	}

	if !PayByCreditIncomeValidation(in) {
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

func (p *PurchaseService) checkPermission(ctx context.Context, cellPhoneNumber string, purchaseType profileProto.PurchasePermissionRequest_PurchaseType) *proto.PurchaseResponse {
	profileResponse, err := p.profileClient.PurchasePermission(ctx, &profileProto.PurchasePermissionRequest{
		UserCellPhone: cellPhoneNumber,
		PurchaseType:  purchaseType,
	})
	if err != nil {
		// log err
		return &proto.PurchaseResponse{
			UserPhoneNo: cellPhoneNumber,
			Error:       proto.PurchaseResponse_Unknown,
		}
	}

	if !profileResponse.Permission {
		return &proto.PurchaseResponse{
			UserPhoneNo: cellPhoneNumber,
			Error:       proto.PurchaseResponse_Permission,
		}
	}

	return nil
}
