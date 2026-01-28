package constants

import "compay/common/types"

const (
	FeeTypeOnlyWithdraw types.FeeType = "OnlyWithdraw"
	FeeTypeFull         types.FeeType = "Full"
)

var FeeTypes = []types.FeeType{FeeTypeOnlyWithdraw, FeeTypeFull}
