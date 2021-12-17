package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres drivers

	"github.com/JasurbekUz/ToDo-service/config"
)

func ConnectionToDB(cfg config.Config) (*sqlx.DB, error) {
	psqlString := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	)

	connDb, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		return nil, err
	}

	return connDb, nil
}

func ConnectDBForSuite(cfg config.Config) (*sqlx.DB, func()) {
	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	connDb, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		panic(err)
	}

	cleanUpFunc := func() {
		connDb.Close()
	}

	return connDb, cleanUpFunc
}
