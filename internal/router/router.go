package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kholodmv/GophKeeper/internal/handlers"
	mw "github.com/kholodmv/GophKeeper/internal/middleware/auth"
)

func Router(h *handlers.Handler) *gin.Engine {
	r := gin.Default()

	auth := r.Group("/")
	{
		auth.Use(mw.Auth())
		auth.POST("/api/secret/create", h.CreateSecret)
		auth.POST("/api/secret/read", h.ReadSecret)
	}
	r.POST("/api/user/register", h.RegisterUser)
	r.POST("/api/user/login", h.LoginUser)

	return r
}
