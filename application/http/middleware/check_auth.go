package middleware

// func CheckAuth(args ...bool) gin.HandlerFunc {
// 	return func(context *gin.Context) {
// 		err := esjwt.ValidateJWT(context)
// 		if err != nil {
// 			helpers.GinErrorResponse(context, err)
// 			context.Abort()
// 			return
// 		}
// 		if len(args) > 0 && args[0] {
// 			context.Next()
// 			return
// 		}
// 		CheckTotp(context)
// 	}
// }

// func CheckTotp(context *gin.Context) {
// 	if err := estotp.ValidateRequest(context); err != nil {
// 		helpers.GinErrorResponse(context, err)
// 		context.Abort()
// 		return
// 	}
// 	context.Next()
// }
