package controller

import (
	"aydin-tracker/aydinins-backend/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type UpdateProductRequestBody struct {
	Name         string  `json:"name"`
	Barcode      string  `json:"barcode"`
	Image        string  `json:"image"`
	Stock        int     `json:"stock"`
	SuppliedFrom string  `json:"suppliedfrom"`
	BuyPrice     float32 `json:"buyprice"`
	SellPrice    float32 `json:"sellprice"`
}

func (h *Handler) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	key := ctx.Param("key")

	if key != viper.Get("KEY").(string) {
		ctx.AbortWithStatus(http.StatusBadRequest) // server will not continue to process due to wrong key value
		return
	}

	body := UpdateProductRequestBody{}
	var product models.Product

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if result := h.DB.First(&product, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	product.Barcode = body.Barcode
	product.BuyPrice = body.BuyPrice
	product.Image = body.Image
	product.Name = body.Name
	product.SellPrice = body.SellPrice
	product.Stock = body.Stock
	product.SuppliedFrom = body.SuppliedFrom

	h.DB.Save(&product)

	ctx.JSON(http.StatusOK, &product)

}
