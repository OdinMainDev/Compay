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

func getUsdToRubRateFromCbr() (float64, error) {
	body, err := doRequest(parserCfg["cbr"])
	if err != nil {
		return 0, fmt.Errorf("cbr request error: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(body))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "windows-1251":
			return transform.NewReader(input, charmap.Windows1251.NewDecoder()), nil
		default:
			return nil, fmt.Errorf("unsupported charset: %s", charset)
		}
	}

	var valCurs ValCurs
	if err := decoder.Decode(&valCurs); err != nil {
		return 0, fmt.Errorf("cbr xml parse error: %w", err)
	}
	for _, valute := range valCurs.Valutes {
		if strings.ToUpper(valute.CharCode) == "USD" {

			valueStr := strings.ReplaceAll(valute.Value, ",", ".")
			rate, err := strconv.ParseFloat(valueStr, 64)
			if err != nil {
				return 0, fmt.Errorf("cbr rate parse error: %w", err)
			}
			return rate, nil
		}
	}
	return 0, fmt.Errorf("USD not found in CBR data")
}
