package main

import (
	"fmt"
	"log"
	"product_service/api"
	"product_service/config"
	"product_service/internal/service"
	"product_service/pkg"
)

func main() {
	cfg := config.Load("./config.yaml")
	db, err := pkg.InitDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	productSvc := service.NewProductService(db) 

	copyService := pkg.NewCopyService(productSvc) 
	go func() {
		if err := copyService.Run(cfg); err != nil {
			log.Fatal(err)
		}
	}()

	r := api.NewGin(db)
	addr := fmt.Sprintf(":%s", cfg.ServicePort)
	r.Run(addr)
}
