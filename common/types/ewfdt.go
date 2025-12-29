package types

import "efw/services/efwdt"

const (
	// BASE TYPES
	EFWDT_FilterTypeInteger efwdt.FilterType = "integer"
	EFWDT_FilterTypeNumber  efwdt.FilterType = "number"
	EFWDT_FilterTypeText    efwdt.FilterType = "text"
	EFWDT_FilterTypeBoolean efwdt.FilterType = "boolean"
	EFWDT_FilterTypeList    efwdt.FilterType = "list"
	EFWDT_FilterTypeDate    efwdt.FilterType = "date"

	// TYPES WITH STATIC DATA

	// TYPES WITH QUERY
	EFWDT_FilterTypeMerchantId efwdt.FilterType = "merchant_id"
)
