package service

import (
	"context"
	"database/sql"
	"product_service/internal/storage"
	"product_service/proto"
)

type ProductService struct {
	proto.UnimplementedProductServiceServer
	ProductRepo storage.ProductRepo
}

func NewProductService(db *sql.DB) *ProductService {
	return &ProductService{
		ProductRepo: *storage.NewProductRepo(db),
	}
}

func (s *ProductService) GetProduct(ctx context.Context, req *proto.ProductRequest) (*proto.Product, error) {
	product, err := s.ProductRepo.GetProduct(req)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) CreateProduct(ctx context.Context, req *proto.Product) (*proto.ProductResponse, error) {
	resp, err := s.ProductRepo.CreateProduct(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, req *proto.Product) (*proto.ProductResponse, error) {
	resp, err := s.ProductRepo.UpdateProduct(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, req *proto.ProductRequest) (*proto.ProductResponse, error) {
	resp, err := s.ProductRepo.DeleteProduct(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
