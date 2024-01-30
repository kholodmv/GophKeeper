package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kholodmv/GophKeeper/internal/models"
	"io"
	"net/http"
)

// CreateSecret - create secret
func (h *Handler) CreateSecret(gctx *gin.Context) {
	var (
		secret *models.Secret
	)

	body, err := io.ReadAll(gctx.Request.Body)
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, err.Error())
	}

	err = json.Unmarshal(body, &secret)
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	userID, _ := gctx.Get("userID")
	secret.UserID = userID.(uint)

	err = h.KeeperService.CreateSecret(secret)
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	gctx.JSON(http.StatusCreated, gin.H{"message": "the new secret has been created!"})
}
