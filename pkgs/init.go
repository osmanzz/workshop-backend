package pkgs

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", "postgres://workshop:workshop@localhost:5001/workshop?sslmode=disable")
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	ping(db)

	return db, nil
}

func ping(db *sqlx.DB) {
	if err := db.Ping(); err != nil {
		log.Panic(err)
	}
}
