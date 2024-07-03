package db

import (
	"fmt"
	"user_service/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" 
)

func ConnectToDB(cfg config.Config) (*sqlx.DB, error, func()) {
	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)

	connDb, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		return nil, err, nil
	}

	cleanUpFunc := func() {
		connDb.Close()
	}

	return connDb, nil, cleanUpFunc
}

func ConnectToDBForSuite(cfg config.Config) (*sqlx.DB, func(), error) {
	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)
	connDB, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		return nil, nil, err
	}
	cleanUpFunc := func() {
		connDB.Close()
	}
	return connDB, cleanUpFunc, nil
}
