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

	routes.GET("/", h.GetProduct)
}
