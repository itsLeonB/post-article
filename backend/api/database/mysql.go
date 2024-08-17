package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectMysql() (*sql.DB, error) {
	url := os.Getenv("DB_URL")

	db, err := sql.Open("mysql", url)
	if err != nil {
		log.Printf("error on ConnectMysql(): %s\ntype: %T\ndetails: %v\n", err.Error(), err, err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("error on ConnectMysql(): %s\ntype: %T\ndetails: %v\n", err.Error(), err, err)
		return nil, err
	}

	return db, nil
}
