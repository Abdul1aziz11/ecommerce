package api

import (
	"database/sql"
	handlers "product_service/api/handler"
	"product_service/internal/storage"

	"github.com/gin-gonic/gin"
)

func NewGin(db *sql.DB) *gin.Engine {
	r := gin.Default()

	productRepo := storage.NewProductRepo(db)
	handler := handlers.NewHandler(productRepo)

	r.POST("/product/create", handler.CreateProduct)
	r.GET("/product/:product_id", handler.GetProduct)
	r.PUT("/product/update/:product_id", handler.UpdateProduct)
	r.DELETE("/product/delete/:product_id", handler.DeleteProduct)

	return r
}
