package cperrors

import "fmt"

var ErrorApiV1BadCurrency = fmt.Errorf("%s", "ApiV1BadCurrency")
var ErrorApiV1BadMerchantToken = fmt.Errorf("%s", "ApiV1BadMerchantToken")
var ErrorApiV1InvalidSign = fmt.Errorf("%s", "ApiV1InvalidSign")
