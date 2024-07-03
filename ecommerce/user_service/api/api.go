package api

import (
	"database/sql"
	"user_service/api/handler"
	"user_service/internal/storage"

	"github.com/gin-gonic/gin"
)

func NewGin(db *sql.DB) *gin.Engine {
	router := gin.Default()

	userRepo := storage.NewUserRepo(db)
	handler := handlers.NewHandler(userRepo)

	router.GET("/user/:user_id", handler.GetUser)
	router.POST("/user", handler.CreateUser)
	router.DELETE("/user/:user_id", handler.DeleteUser)
	router.PUT("/user/:user_id", handler.UpdateUser)

	return router
}
