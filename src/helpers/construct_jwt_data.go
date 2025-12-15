package helpers

import (
	"compay/common/constants"
	"compay/src/models/user"
	"efw/core/efwjwt"
)

func ConstructJwtData(userModel *user.Model) []efwjwt.JwtData {
	jwtData := []efwjwt.JwtData{}
	jwtData = append(jwtData, efwjwt.JwtData{
		Key:   constants.JwtClaimsKeyIsAdmin,
		Value: userModel.MerchantId == 0,
	})
	jwtData = append(jwtData, efwjwt.JwtData{
		Key:              constants.JwtClaimsKeyUserId,
		Value:            userModel.GetIdInt64(),
		IsEncryptedValue: true,
	})
	jwtData = append(jwtData, efwjwt.JwtData{
		Key:              constants.JwtClaimsKeyMerchantId,
		Value:            userModel.MerchantId,
		IsEncryptedValue: true,
	})

	return jwtData
}
