package handlers

import (
	"log"
	"net/http"
	"os"
	"product_service/internal/storage"
	"product_service/proto"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	ProductRepo storage.ProductRepo
	ErrorLog    *log.Logger
	InfoLog     *log.Logger
}

func NewHandler(productRepo storage.ProductRepo) *Handler {
	file, err := os.OpenFile("./logging/product_service.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	errLog := log.New(file, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(file, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	return &Handler{ProductRepo: productRepo, ErrorLog: errLog, InfoLog: infoLog}
}

func (h *Handler) CreateProduct(ctx *gin.Context) {
	var req proto.Product
	if err := ctx.BindJSON(&req); err != nil {
		h.ErrorLog.Printf("Error binding JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.ProductRepo.CreateProduct(&req)
	if err != nil {
		h.ErrorLog.Printf("Error creating product: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	h.InfoLog.Printf("Product created successfully: %v", resp)
	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateProduct(ctx *gin.Context) {
	productID := ctx.Param("product_id")
	if productID == "" {
		h.ErrorLog.Println("product_id is required")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "product_id is required"})
		return
	}

	var req proto.Product
	if err := ctx.BindJSON(&req); err != nil {
		h.ErrorLog.Printf("Error binding JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Id = productID

	resp, err := h.ProductRepo.UpdateProduct(&req)
	if err != nil {
		h.ErrorLog.Printf("Error updating product: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	h.InfoLog.Printf("Product updated successfully: %v", resp)
	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteProduct(ctx *gin.Context) {
	productID := ctx.Param("product_id")
	if productID == "" {
		h.ErrorLog.Println("product_id is required")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "product_id is required"})
		return
	}

	resp, err := h.ProductRepo.DeleteProduct(&proto.ProductRequest{Id: productID})
	if err != nil {
		h.ErrorLog.Printf("Error deleting product: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	h.InfoLog.Printf("Product deleted successfully: %v", resp)
	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetProduct(ctx *gin.Context) {
	productID := ctx.Param("product_id")
	if productID == "" {
		h.ErrorLog.Println("product_id is required")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "product_id is required"})
		return
	}

	resp, err := h.ProductRepo.GetProduct(&proto.ProductRequest{Id: productID})
	if err != nil {
		h.ErrorLog.Printf("Error getting product: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get product"})
		return
	}

	h.InfoLog.Printf("Product retrieved successfully: %v", resp)
	ctx.JSON(http.StatusOK, resp)
}
