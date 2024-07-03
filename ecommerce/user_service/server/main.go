package main

import (
	"fmt"
	"log"
	"user_service/api"
	"user_service/config"
	"user_service/internal/service"
	"user_service/pkg"
)

func main() {
	cfg := config.Load("./")

	db, err := pkg.InitDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	defer db.Close()

	userService := service.UserService(db)

	grpcServer := pkg.NewCopyService(userService)
	go func() {
		if err := grpcServer.RUN(cfg); err != nil {
			log.Fatal("Failed to start gRPC server:", err)
		}
	}()

	httpServer := api.NewGin(db)
	httpAddr := fmt.Sprintf(":%s", cfg.ServicePort)
	httpServer.Run(httpAddr)
}
