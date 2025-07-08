package wallet

type requestAuth struct {
	PinCode string `json:"pin_code" binding:"required,pinCode"`
	Seed    string `json:"seed" binding:"omitempty,seedPhrase"`
}
