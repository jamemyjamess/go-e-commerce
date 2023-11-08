package postgres

import (
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jamemyjamess/go-e-commerce/config"
	"github.com/jmoiron/sqlx"
)

func ConnectPostgresDB(cfg config.IDbConfig) *sqlx.DB {
	db, err := sqlx.Connect("pgx", cfg.Url())
	if err != nil {
		log.Fatalf("connect to db failed: %v\n", err)
	}
	db.DB.SetMaxOpenConns(cfg.MaxOpenConns())
	return db
}
