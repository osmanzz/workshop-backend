package pkgs

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func InitDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", "root:@tcp(localhost:3306)/workshop")
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
