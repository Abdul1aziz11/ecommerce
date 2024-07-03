package pkg

import (
	"database/sql"
	"fmt"
	"product_service/config"

	_ "github.com/lib/pq"
)

func InitDB(cfg config.Config) (*sql.DB, error) {
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Database)

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return db, nil
}
