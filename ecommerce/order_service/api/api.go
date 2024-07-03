package api

import (
	"database/sql"
	"order_service/api/handler"
	"order_service/internal/storage"

	"github.com/gin-gonic/gin"
)

func NewGin(db *sql.DB) *gin.Engine {
	r := gin.Default()

	orderHandler := handler.NewHandler(
		storage.NewOrderRepo(db),
	)

	r.GET("/order/:order_id", orderHandler.GetOrder)
	r.POST("/create", orderHandler.CreateOrder)
	r.PUT("/update/:order_id", orderHandler.UpdateOrder)
	r.DELETE("/delete/:order_id", orderHandler.DeleteOrder)

	return r
}
