package rates

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
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

var cbrRubVal = 0.
var cbrRubCnt = 0

func getUsdToRubRateFromCbr() (rV float64, rE error) {
	if cbrRubVal > 0 {
		if cbrRubCnt < 120 {
			cbrRubCnt++
			return cbrRubVal, nil
		} else {
			cbrRubCnt = 0
			cbrRubVal = 0
		}
	}

	defer func() { cbrRubVal = rV }()
	body, err := doRequest(parserCfg["cbr"])
	if err != nil {
		return 0, fmt.Errorf("cbr 1: cbr request error: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(body))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "windows-1251":
			return transform.NewReader(input, charmap.Windows1251.NewDecoder()), nil
		default:
			return nil, fmt.Errorf("cbr 2: unsupported charset: %s", charset)
		}
	}

	var valCurs ValCurs
	if err := decoder.Decode(&valCurs); err != nil {
		return 0, fmt.Errorf("cbr 3: cbr xml parse error: %w", err)
	}
	for _, valute := range valCurs.Valutes {
		if strings.ToUpper(valute.CharCode) == "USD" {

			valueStr := strings.ReplaceAll(valute.Value, ",", ".")
			rate, err := strconv.ParseFloat(valueStr, 64)
			if err != nil {
				return 0, fmt.Errorf("cbr 4: cbr rate parse error: %w", err)
			}
			return rate, nil
		}
	}
	return 0, fmt.Errorf("cbr 5: USD not found in CBR data")
}
