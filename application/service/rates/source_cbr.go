package rates

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type ValCurs struct {
	// XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	// ID       string `xml:"ID,attr"`
	// NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	// Nominal  int    `xml:"Nominal"`
	// Name     string `xml:"Name"`
	Value string `xml:"Value"`
}

func getUsdToRubRateFromCbr() (float64, error) {
	body, err := doRequest(parserCfg["cbr"])
	if err != nil {
		return 0, fmt.Errorf("cbr request error: %w", err)
	}

	var valCurs ValCurs
	if err := xml.Unmarshal(body, &valCurs); err != nil {
		return 0, fmt.Errorf("cbr xml parse error: %w", err)
	}
	for _, valute := range valCurs.Valutes {
		if strings.ToUpper(valute.CharCode) == "USD" {

			valueStr := strings.Replace(valute.Value, ",", ".", -1)
			rate, err := strconv.ParseFloat(valueStr, 64)
			if err != nil {
				return 0, fmt.Errorf("cbr rate parse error: %w", err)
			}
			return rate, nil
		}
	}
	return 0, fmt.Errorf("USD not found in CBR data")
}
