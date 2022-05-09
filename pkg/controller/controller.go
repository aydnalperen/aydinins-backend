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
	routes.POST("/", h.AddProduct)
	routes.POST("/:id", h.DeleteProduct)
	routes.POST("/:id", h.UpdateProduct)
}
