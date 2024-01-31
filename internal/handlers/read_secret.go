package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	req struct {
		Title string `json:"title"`
	}
)

// ReadSecret - read secret by name
func (h *Handler) ReadSecret(ctx *gin.Context) {
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	userID, _ := ctx.Get("userID")
	secret, err := h.KeeperService.ReadSecret(req.Title, userID.(uint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	response, err := json.Marshal(secret)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.Data(http.StatusOK, "application/json", response)
}
