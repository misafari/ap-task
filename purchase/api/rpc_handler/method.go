package rpc_handler

import (
	profileProto "asanpardakht/profile/api/proto"
	"asanpardakht/purchase/api/proto"
	"context"
)

func (p *PurchaseService) ByCartProcess(ctx context.Context) <-chan *proto.PurchaseResponse {
	ch := make(chan *proto.PurchaseResponse)
	payByCartRequest := ctx.Value(PayByCartRequestKey).(*proto.PayByCartRequest)

	//time.Sleep(3 * time.Second)

	go func() {
		defer close(ch)
		permissionDenied := p.checkPermission(ctx, payByCartRequest.GetUserPhoneNo(), profileProto.PurchasePermissionRequest_ByCart)
		if permissionDenied != nil {
			ch <- permissionDenied
			return
		}
		if !PayByCartIncomeValidation(payByCartRequest) {
			ch <- &proto.PurchaseResponse{
				UserPhoneNo: payByCartRequest.GetUserPhoneNo(),
				Error:       proto.PurchaseResponse_Info,
			}
			return
		}

		ch <- &proto.PurchaseResponse{
			UserPhoneNo: payByCartRequest.GetUserPhoneNo(),
			Error:       MakeFakeDecision(),
		}
	}()

	return ch
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
