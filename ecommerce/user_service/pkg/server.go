package pkg

import (
	"fmt"
	"net"
	"user_service/config"
	"user_service/internal/service"
	"user_service/proto"

	"google.golang.org/grpc"
)

type CopyService struct {
	service *service.UserServiceServer
}

func NewCopyService(s *service.UserServiceServer) *CopyService {
	return &CopyService{service: s}
}

func (cs *CopyService) Run(cfg config.Config) error {
	address := fmt.Sprintf(":%s", cfg.UserServicePort)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	proto.RegisterUserServiceServer(server, cs.service)

	return server.Serve(listener)
}
