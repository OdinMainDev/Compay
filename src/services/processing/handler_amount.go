package processing

import (
	"compay/common/cperrors"
	"efw/components/entity"
	"efw/core/efwamount"
)

func (that *Processing) calculateDepositAmount(amountFloat float64) (entity.Amount, error) {
	if that.getMerchantCurrency().MinAmount == 0 && amountFloat == 0 {
		return 0, nil
	} else if amountFloat == 0 {
		return 0, cperrors.ErrorProcessingDepositAmountIsRequired
	}

	amount := efwamount.FloatToAmount(amountFloat, that.getCurrency().MinorUnits)
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
