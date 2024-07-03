package pkg

import (
	"fmt"
	"net"
	"product_service/config"
	"product_service/internal/service"
	"product_service/proto"

	"google.golang.org/grpc"
)

type CopyServer struct {
	service *service.ProductService
}

func NewCopyServer(s *service.ProductService) *CopyServer {
	return &CopyServer{service: s}
}

func (s *CopyServer) Run(cfg config.Config) error {
	address := fmt.Sprintf(":%s", cfg.ProductServicePort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	proto.RegisterProductServiceServer(server, s.service)

	fmt.Printf("gRPC server running on %s\n", address)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}
