package rpc_handler

import (
	"asanpardakht/purchase/api/proto"
	"math/rand"
	"time"
)

func PayByCartIncomeValidation(in *proto.PayByCartRequest) bool {
	return in.GetCartNumber() != "" && in.GetUserPhoneNo() != "" &&
		in.GetAmount() != 0 && CartNumberValidation(in.GetCartNumber())
}

func PayByWalletIncomeValidation(in *proto.PayByWalletRequest) bool {
	return in.GetWalletNumber() != "" && in.GetUserPhoneNo() != "" && in.GetAmount() != 0
}

func PayByCreditIncomeValidation(in *proto.PayByCreditRequest) bool {
	return in.GetCreditId() != "" && in.GetUserPhoneNo() != "" && in.GetAmount() != 0
}

func CartNumberValidation(cardNumber string) bool {
	return len(cardNumber) == 16
}

func MakeFakeDecision() proto.PurchaseResponse_Error {
	source := rand.NewSource(time.Now().UnixNano())
	customRand := rand.New(source)
	randInt := customRand.Int31()

	var e proto.PurchaseResponse_Error
	switch {
	case randInt%3 == 0:
		e = proto.PurchaseResponse_Credit
	case randInt%5 == 0:
		e = proto.PurchaseResponse_Unknown
	default:
		e = proto.PurchaseResponse_None
	}

	return e
}
