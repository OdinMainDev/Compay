package processing

import "compay/common/constants"

type requetsUrl string

const (
	requestCreateWallet requetsUrl = "/generate-wallet"
)

func (that *Processing) getRequestUrl(url requetsUrl) string {
	return that.getRequestHost() + string(url)
}

func (that *Processing) getRequestHost() string {
	switch that.getCurrency().CodeAlpha {
	case constants.BtcCodeAlpha:
		return constants.BtcNodeUrl
	}

	return ""
}
