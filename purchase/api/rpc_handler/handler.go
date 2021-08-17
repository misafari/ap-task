package rpc_handler

import (
	profileProto "asanpardakht/profile/api/proto"
	"asanpardakht/purchase/api/proto"
	"context"
	"time"
)

type PurchaseService struct {
	profileClient profileProto.ProfileClient
	deadline      time.Duration
}

const (
	PayByCartRequestKey = "PayByCartRequestKey"
)

func NewPurchaseService(profileClient profileProto.ProfileClient, deadline time.Duration) *PurchaseService {
	return &PurchaseService{
		deadline:      deadline,
		profileClient: profileClient,
	}
}

func (p *PurchaseService) PayByCart(ctx context.Context, in *proto.PayByCartRequest) (*proto.PurchaseResponse, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(p.deadline))
	ctx = context.WithValue(ctx, PayByCartRequestKey, in)
	defer cancel()

	ch := p.ByCartProcess(ctx)

	select {
	case <-ctx.Done():
		return &proto.PurchaseResponse{
			UserPhoneNo: in.GetUserPhoneNo(),
			Error:       proto.PurchaseResponse_TimeOut,
		}, nil
	case r := <-ch:
		return r, nil
	}
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
