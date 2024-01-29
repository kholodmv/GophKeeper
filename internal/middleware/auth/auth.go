package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/kholodmv/GophKeeper/internal/utils/jwt"
	"net/http"
)

// Auth - authentication check
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		userID, err := jwt.GetUserIDFromToken(token)
		if err != nil || userID == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("userID", userID)
		ctx.Next()
	}
}
