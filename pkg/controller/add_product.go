package controller

import (
	"aydin-tracker/aydinins-backend/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddBookRequestBody struct {
	Name         string  `json:"name"`
	Barcode      string  `json:"barcode"`
	Image        string  `json:"image"`
	Stock        int     `json:"stock"`
	SuppliedFrom string  `json:"suppliedfrom"`
	BuyPrice     float32 `json:"buyprice"`
	SellPrice    float32 `json:"sellprice"`
}

func (h *Handler) AddBook(ctx *gin.Context) {
	body := AddBookRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var product models.Product

	product.Barcode = body.Barcode
	product.BuyPrice = body.BuyPrice
	product.Image = body.Image
	product.Name = body.Name
	product.SellPrice = body.SellPrice
	product.Stock = body.Stock
	product.SuppliedFrom = body.SuppliedFrom

	if result := h.DB.Create(&product); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	ctx.JSON(http.StatusOK, &product)
}
