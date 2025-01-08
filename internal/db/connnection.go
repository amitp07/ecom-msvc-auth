package db

import (
	"database/sql"
	"log"

	"github.com/amitp07/ecom-msvc-auth/internal/env"
	_ "github.com/lib/pq"
)

func InitDb() *sql.DB {

	var db *sql.DB
	var err error
	if db, err = sql.Open("postgres", env.Env.DbConnectionStr); err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
