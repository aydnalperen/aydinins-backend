package controller

import (
	"aydin-tracker/aydinins-backend/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	var product models.Product

	if result := h.DB.First(&product, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&product)
	ctx.Status(http.StatusOK)
}
