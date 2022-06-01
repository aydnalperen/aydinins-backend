package controller

import (
	"aydin-tracker/aydinins-backend/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func (h *Handler) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	key := ctx.Param("key")
	if key != viper.Get("KEY").(string) {
		ctx.AbortWithStatus(http.StatusBadRequest) // server will not continue to process due to wrong key value
		return
	}
	var product models.Product

	if result := h.DB.First(&product, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&product)
	ctx.Status(http.StatusOK)
}
