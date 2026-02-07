package processing

import (
	"compay/common/constants"
	"efw/core/efwlog"
)

type requestUrl string

const (
	requestCreateWallet requestUrl = "/generate-wallet"
	requestWithdraw     requestUrl = "/withdraw"
)

func (that *Processing) getRequestUrl(url requestUrl) string {
	return that.getRequestHost() + string(url)
}

func (that *Processing) getRequestHost() string {
	switch that.getCurrency().CodeAlpha {
	case constants.BtcCodeAlpha:
		return constants.BtcApiUrl

	default:
		efwlog.Panic("BAD CONFIGURATION getRequestHost: ", that.getCurrency().CodeAlpha)
	}

	return ""
}
