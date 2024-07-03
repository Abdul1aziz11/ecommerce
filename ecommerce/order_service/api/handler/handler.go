package handler

import (
	"context"
	"log"
	"net/http"
	"order_service/internal/storage"
	"order_service/proto/order"
	"os"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	OrderRepo storage.OrderRepo
	ErrorLog  *log.Logger
	InfoLog   *log.Logger
}

func NewHandler(orderRepo storage.OrderRepo) *Handler {
	errLog := log.New(os.Stdout, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	return &Handler{OrderRepo: orderRepo, ErrorLog: errLog, InfoLog: infoLog}
}

func (h *Handler) CreateOrder(ctx *gin.Context) {
	var req order.OrderRequest
	if err := ctx.BindJSON(&req); err != nil {
		h.ErrorLog.Printf("Error binding JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.OrderRepo.CreateOrder(context.Background(), &req)
	if err != nil {
		h.ErrorLog.Printf("Error creating order: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.InfoLog.Printf("Order created successfully: %v", resp)
	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateOrder(ctx *gin.Context) {
	var req order.OrderRequest
	if err := ctx.BindJSON(&req); err != nil {
		h.ErrorLog.Printf("Error binding JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Id = ctx.Param("order_id")

	resp, err := h.OrderRepo.UpdateOrder(context.Background(), &req)
	if err != nil {
		h.ErrorLog.Printf("Error updating order: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.InfoLog.Printf("Order updated successfully: %v", resp)
	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteOrder(ctx *gin.Context) {
	orderID := ctx.Param("order_id")
	if orderID == "" {
		h.ErrorLog.Println("order_id is required")
		ctx.JSON(http.StatusBadRequest, "order_id is required")
		return
	}

	resp, err := h.OrderRepo.DeleteOrder(context.Background(), &order.OrderRequest{Id: orderID})
	if err != nil {
		h.ErrorLog.Printf("Error deleting order: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.InfoLog.Printf("Order deleted successfully: %v", resp)
	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetOrder(ctx *gin.Context) {
	orderID := ctx.Param("order_id")
	if orderID == "" {
		h.ErrorLog.Println("order_id is required")
		ctx.JSON(http.StatusBadRequest, "order_id is required")
		return
	}

	resp, err := h.OrderRepo.GetOrder(context.Background(), &order.OrderRequest{Id: orderID})
	if err != nil {
		h.ErrorLog.Printf("Error getting order: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.InfoLog.Printf("Order retrieved successfully: %v", resp)
	ctx.JSON(http.StatusOK, resp)
}
