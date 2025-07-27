package cperrors

import "fmt"

var ErrorAdminCurrencyCodeAlphaDuplicate = fmt.Errorf("%s", "CodeAlphaDuplicate")
var ErrorAdminCurrenciesImportAlreadyStarted = fmt.Errorf("%s", "CurrenciesImportAlreadyStarted")
var ErrorAdminMerchantTokenMerchantDuplicate = fmt.Errorf("%s", "TokenDuplicate")
