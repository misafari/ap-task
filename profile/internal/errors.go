package internal

import "errors"

var (
	CellPhoneIsRequired = errors.New("user cell phone is required")
	PurchaseTypeIsRequired = errors.New("purchase type is required")
)
