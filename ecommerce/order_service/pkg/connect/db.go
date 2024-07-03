package pkg

import (
	"database/sql"
	"fmt"
	"order_service/config"

	_ "github.com/lib/pq"
)

func InitDB(cfg config.Config) (*sql.DB, error) {
	dbData := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Database,
	)

	db, err := sql.Open("postgres", dbData)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
