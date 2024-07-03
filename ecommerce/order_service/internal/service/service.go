package service

import (
	"context"
	"order_service/proto/product"
	"order_service/proto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func UserService(userId string) (*user.User, error) {
	conn, err := grpc.Dial("localhost:8001", grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	userGR := user.NewUserServiceClient(conn)
	resp, err := userGR.GetUser(context.Background(), &user.UserRequest{Id: userId})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func ProductService(productId string) (*product.Product, error) {
	conn, err := grpc.Dial("localhost:8002", grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	productGR := product.NewProductServiceClient(conn)
	resp, err := productGR.GetProduct(context.Background(), &product.ProductRequest{Id: productId})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
