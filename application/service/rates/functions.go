package rates

import (
	"compay/core/esrequest"
)

func doRequest(url string) ([]byte, error) {
	resp, _, err := esrequest.Request(esrequest.RequestHelper{
		Url:    url,
		Method: "GET",
		Headers: map[string]string{
			"Accept-Language": "ru-RU,ru;q=0.8,en-US;q=0.6,en;q=0.4",
			"Cache-Control":   "no-cache",
			"Pragma":          "no-cache",
			"User-Agent":      "Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6400.0 Safari/537.36",
		},
		DisableLog: true,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
