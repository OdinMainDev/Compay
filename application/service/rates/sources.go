package rates

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

type KucoinResp struct {
	Code string `json:"code"`
	Data struct {
		Ticker []struct {
			Symbol string `json:"symbol"`
			Last   string `json:"last"`
		} `json:"ticker"`
	} `json:"data"`
}

func getKucoin() (map[string]map[string]float64, error) {
	body, err := doRequest(parserCfg["kucoin"])
	if err != nil {
		return nil, err
	}
	var ku KucoinResp
	if err := json.Unmarshal(body, &ku); err != nil {
		return nil, err
	}
	r := map[string]map[string]float64{}
	for _, t := range ku.Data.Ticker {
		pair := strings.Split(t.Symbol, "-")
		if len(pair) != 2 {
			continue
		}
		base, quote := pair[0], pair[1]
		if !inList(base) || !inList(quote) {
			continue
		}
		val, err := strconv.ParseFloat(t.Last, 64)
		if err != nil {
			continue
		}
		if _, ok := r[base]; !ok {
			r[base] = map[string]float64{}
		}
		r[base][quote] = val
	}
	return r, nil
}

func getRapira() (float64, error) {
	body, err := doRequest(parserCfg["rapira"])
	if err != nil {
		return 0, err
	}
	var data []map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return 0, err
	}
	for _, d := range data {
		if d["symbol"] == "USDT/RUB" {
			if val, ok := d["high"].(float64); ok {
				return val, nil
			}
		}
	}
	return 0, errors.New("RUB not found")
}

type CBRValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []struct {
		CharCode string `xml:"CharCode"`
		Value    string `xml:"Value"`
	} `xml:"Valute"`
}

func getCbrEUR() (float64, error) {
	body, err := doRequest(parserCfg["cbr"])
	if err != nil {
		return 0, err
	}

	decoder := xml.NewDecoder(strings.NewReader(string(body)))
	decoder.CharsetReader = func(charset2 string, input io.Reader) (io.Reader, error) {
		return charset.NewReaderLabel(charset2, input)
	}

	var valCurs CBRValCurs
	if err := decoder.Decode(&valCurs); err != nil {
		return 0, err
	}

	for _, val := range valCurs.Valutes {
		if strings.ToUpper(val.CharCode) == "EUR" {
			valStr := strings.Replace(val.Value, ",", ".", 1)
			return strconv.ParseFloat(valStr, 64)
		}
	}
	return 0, errors.New("EUR not found")
}
