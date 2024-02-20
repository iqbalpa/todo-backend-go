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