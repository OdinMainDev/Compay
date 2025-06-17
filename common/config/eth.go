package config

var ETHMinorUnits = 18

var ETHContractData = map[string]struct {
	Address string
	Units   int
}{
	"USDT":  {"0xdac17f958d2ee523a2206206994597c13d831ec7", 6},
	"USDC":  {"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", 6},
	"DAI":   {"0x6b175474e89094c44da98b954eedeac495271d0f", 18},
	"CRV":   {"0xD533a949740bb3306d119CC777fa900bA034cd52", 18},
	"MATIC": {"0x7d1afa7b718fb893db30a3abc0cfc608aacfebb0", 18},
	"LINK":  {"0x514910771AF9Ca656af840dff83E8264EcF986CA", 18},
	"WBTC":  {"0x2260fac5e5542a773aa44fbcfedf7c193bc2c599", 8},
	"UNI":   {"0x1f9840a85d5af5bf1d1762f925bdaddc4201f984", 18},
	"AAVE":  {"0x7fc66500c84a76ad7e9c93437bfc5ac33e2ddae9", 18},
	"SHIB":  {"0x95aD61b0a150d79219dCF64E1E6Cc01f0B64C4cE", 18},
	"BNB":   {"0xB8c77482e45F1F44dE1745F52C74426C631bDD52", 18},
	"ARB":   {"0xB50721BCf8d664c30412Cfbc6cf7a15145234ad1", 18},
}
