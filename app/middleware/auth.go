package middleware

import (
	"main/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get token from header
		token := ctx.Request.Header.Get("Authorization")
		err := utils.VerifyToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{
					"error": err.Error(),
				},
			)
			return
		}
		ctx.Set("token", token)
		ctx.Set("isAuthenticated", true)
		ctx.Next()
	}
}

// func UserIdExtractor() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		// check whether the user logged in or not
// 		isAuthenticated := ctx.GetBool("isAuthenticated")
// 		if isAuthenticated {
// 			token := ctx.Request.Header.Get("Authorization")
// 			userId, err := utils.ExtractClaimsUserId(token)
// 			if err != nil {
// 				ctx.AbortWithStatusJSON(http.StatusBadRequest,
// 					gin.H{
// 						"error": err.Error(),
// 					},
// 				)
// 				return
// 			}
// 			ctx.Set("userId", userId)
// 		}
// 		ctx.Next()
// 	}
// }