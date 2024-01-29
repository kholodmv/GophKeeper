package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/kholodmv/GophKeeper/internal/models"
	"net/http"
)

// LoginUser Login - user authorization
func (h *Handler) LoginUser(ctx *gin.Context) {
	var auth *models.Auth
	if err := ctx.ShouldBindJSON(&auth); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	token, err := h.AuthService.Login(auth)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.Writer.Header().Set("Authorization", token)
	ctx.JSON(http.StatusCreated, gin.H{"message": "the user has logged in"})
}
