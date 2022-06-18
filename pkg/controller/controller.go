package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func RouteListener(r *gin.Engine, db *gorm.DB) {
	h := &Handler{
		DB: db,
	}

	routes := r.Group("/products")

	routes.GET("/:id", h.GetProduct)
	routes.GET("/", h.GetProducts)
	routes.POST("/add/:key", h.AddProduct)
	routes.POST("/delete/:id/:key", h.DeleteProduct)
	routes.POST("/update/:id/:key", h.UpdateProduct)
	routes.POST("/images/", h.UploadImage)

	routes.Static("/images", "./images")
}
