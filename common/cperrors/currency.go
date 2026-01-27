package cperrors

import "fmt"

var ErrorCurrencyImportIsAlreadyStarted = fmt.Errorf("%s", "CurrencyImportIsAlreadyStarted")
var ErrorCurrencyCodeAlphaIsTaken = fmt.Errorf("%s", "CurrencyCodeAlphaIsTaken")
