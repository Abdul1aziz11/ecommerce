package main

import (
	"context"
	"log"
	"net/http"

	"order_service/proto/product"
	"order_service/proto/user"

	"google.golang.org/grpc"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	userConn, err := grpc.Dial("localhost:8001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to user service: %v", err)
	}
	defer userConn.Close()
	userClient := user.NewUserServiceClient(userConn)

	productConn, err := grpc.Dial("localhost:8002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to product service: %v", err)
	}
	defer productConn.Close()
	productClient := product.NewProductServiceClient(productConn)

	orderConn, err := grpc.Dial("localhost:8003", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to order service: %v", err)
	}
	defer orderConn.Close()
	orderClient := order.NewOrderServiceClient(orderConn)

	r.Any("/users/*proxyPath", reverseProxy(userClient, "/users"))
	r.Any("/products/*proxyPath", reverseProxy(productClient, "/products"))
	r.Any("/orders/*proxyPath", reverseProxy(orderClient, "/orders"))

	log.Println("API Gateway running on port 8080")
	log.Fatal(r.Run(":8080"))
}

func reverseProxy(client interface{}, prefix string) gin.HandlerFunc {
	return func(c *gin.Context) {
		switch client := client.(type) {
		case user.UserServiceClient:
			resp, err := client.GetUser(context.Background(), &user.UserRequest{Id: c.Param("proxyPath")})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, resp)
		case product.ProductServiceClient:
			resp, err := client.GetProduct(context.Background(), &product.ProductRequest{Id: c.Param("proxyPath")})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, resp)
		case order.OrderServiceClient:
			resp, err := client.GetOrder(context.Background(), &order.OrderRequest{Id: c.Param("proxyPath")})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, resp)
		default:
			c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		}
	}
}
