package rpc_handler

import (
	"asanpardakht/profile/api/proto"
	"asanpardakht/profile/internal"
	"context"
	"math/rand"
	"time"
)

type ProfileServer struct {}

func (p *ProfileServer) PurchasePermission(_ context.Context, request *proto.PurchasePermissionRequest) (*proto.PurchasePermissionResponse, error) {
	if request.GetUserCellPhone() == "" {
		return nil, internal.CellPhoneIsRequired
	}

	if request.GetPurchaseType() == proto.PurchasePermissionRequest_None {
		return nil, internal.PurchaseTypeIsRequired
	}

	source := rand.NewSource(time.Now().UnixNano())
	customRand := rand.New(source)

	per := customRand.Int31() % 2 == 0 || customRand.Int31() % 3 == 0 || customRand.Int31() % 5 == 0

	return &proto.PurchasePermissionResponse{
		Permission: per,
	}, nil
}


func NewProfileServer() *ProfileServer {
	return &ProfileServer{}
}