package service

import (
	"context"
	"database/sql"
	"user_service/internal/storage"
	"user_service/proto"
)

type UserServiceServer struct {
	proto.UnimplementedUserServiceServer
	UserRepo storage.UserRepo
}

func NewUserService(db *sql.DB) *UserServiceServer {
	return &UserServiceServer{UserRepo: *storage.NewUserRepo(db)}
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *proto.UserRequest) (*proto.User, error) {
	user, err := s.UserRepo.GetUser(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}
