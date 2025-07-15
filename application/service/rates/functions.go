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
			"accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
			"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36",

			"sec-ch-ua":                 `"Chromium";v="137", "Not/A)Brand";v="24"`,
			"sec-ch-ua-mobile":          "?0",
			"sec-ch-ua-platform":        `"macOS"`,
			"sec-fetch-dest":            "document",
			"sec-fetch-mode":            "navigate",
			"sec-fetch-site":            "none",
			"sec-fetch-user":            "?1",
			"upgrade-insecure-requests": "1",
		},
		// DisableLog: true,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
