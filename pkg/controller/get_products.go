package controller

import (
	"aydin-tracker/aydinins-backend/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetProducts(ctx *gin.Context) {
	var products []models.Product

	if result := h.DB.Find(&products); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &products)
}
