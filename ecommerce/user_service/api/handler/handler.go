package handlers

import (
	"log"
	"net/http"
	"os"
	"user_service/internal/storage"
	"user_service/proto"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserRepo *storage.UserRepo
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

func NewHandler(userRepo *storage.UserRepo) *Handler {
	file, err := os.OpenFile("./logging/user_service.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	errLog := log.New(file, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(file, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	return &Handler{UserRepo: userRepo, ErrorLog: errLog, InfoLog: infoLog}
}

func (h *Handler) CreateUser(ctx *gin.Context) {
	var user proto.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		h.ErrorLog.Printf("Error binding JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.UserRepo.CreateUser(&user)
	if err != nil {
		h.ErrorLog.Printf("Error creating user: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	h.InfoLog.Printf("User created successfully: %v", resp)
	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateUser(ctx *gin.Context) {
	var user proto.User
	userID := ctx.Param("user_id")
	if userID == "" {
		h.ErrorLog.Println("user_id is required")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		h.ErrorLog.Printf("Error binding JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Id = userID
	resp, err := h.UserRepo.UpdateUser(&user)
	if err != nil {
		h.ErrorLog.Printf("Error updating user: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	h.InfoLog.Printf("User updated successfully: %v", resp)
	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteUser(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	if userID == "" {
		h.ErrorLog.Println("user_id is required")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}
	resp, err := h.UserRepo.DeleteUser(&proto.UserRequest{Id: userID})
	if err != nil {
		h.ErrorLog.Printf("Error deleting user: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	h.InfoLog.Printf("User deleted successfully: %v", resp)
	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetUser(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	if userID == "" {
		h.ErrorLog.Println("user_id is required")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}
	resp, err := h.UserRepo.GetUser(&proto.UserRequest{Id: userID})
	if err != nil {
		h.ErrorLog.Printf("Error getting user: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	h.InfoLog.Printf("User retrieved successfully: %v", resp)
	ctx.JSON(http.StatusOK, resp)
}
