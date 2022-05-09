package controller

import (
	"aydin-tracker/aydinins-backend/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	var product models.Product

	if result := h.DB.First(&product, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	ctx.JSON(http.StatusOK, &product)
}
