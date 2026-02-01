package processing

import (
	"compay/common/cperrors"
	"compay/common/types"
	"compay/src/models/currency"
)

func (that *Processing) calculateDepositAmount(amountFloat float64) (types.Amount, error) {
	if that.getMerchantCurrency().MinAmount == 0 && amountFloat == 0 {
		return 0, nil
	} else if amountFloat == 0 {
		return 0, cperrors.ErrorProcessingDepositAmountIsRequired
	}

	amount := currency.FloatToAmount(amountFloat, that.getCurrency().MinorUnits)
	if that.getMerchantCurrency().MinAmount != 0 {
		if amount < that.getMerchantCurrency().MinAmount {
			return 0, cperrors.ErrorProcessingDepositAmountLessThanMin
		}
	}

	if that.getMerchantCurrency().MaxAmount != 0 {
		if amount > that.getMerchantCurrency().MaxAmount {
			return 0, cperrors.ErrorProcessingDepositAmountBiggerThanMax
		}
	}

	return amount, nil
}
