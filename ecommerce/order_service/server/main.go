package main

import (
	"fmt"
	"log"
	"order_service/api"
	"order_service/config"
	"order_service/pkg"
)

func main() {
	cfg, err := config.Load("./")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := pkg.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	router := api.NewGin(db)

	addr := fmt.Sprintf(":%s", cfg.ServicePort)
	err = router.Run(addr)
	if err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
