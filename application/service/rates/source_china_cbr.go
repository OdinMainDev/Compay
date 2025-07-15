package rates

import (
	"encoding/xml"
	"fmt"
)

type Chart struct {
	XMLName xml.Name `xml:"chart"`
	XAxis   XAxis    `xml:"xaxis"`
	Graphs  Graphs   `xml:"graphs"`
}

type XAxis struct {
	Values []XValue `xml:"value"`
}

type XValue struct {
	XID  int    `xml:"xid,attr"`
	Date string `xml:",chardata"`
}

type Graphs struct {
	Graph []Graph `xml:"graph"`
}

type Graph struct {
	Values []GValue `xml:"value"`
}

type GValue struct {
	XID  int     `xml:"xid,attr"`
	Rate float64 `xml:",chardata"`
}

func getUsdToCnyRateFromChCbr() (rate float64, err error) {
	body, err := doRequest("https://www.chinamoney.com.cn/r/cms/www/chinamoney/data/fx/ccpr-chrt-usd.xml")
	if err != nil {
		return 0, fmt.Errorf("request error: %w", err)
	}

	var chart Chart
	if err := xml.Unmarshal(body, &chart); err != nil {
		return 0, fmt.Errorf("xml parse error: %w", err)
	}

	if len(chart.XAxis.Values) == 0 || len(chart.Graphs.Graph) == 0 {
		return 0, fmt.Errorf("invalid XML structure")
	}

	lastIdx := len(chart.XAxis.Values) - 1
	rate = chart.Graphs.Graph[0].Values[lastIdx].Rate

	return rate, nil
}
