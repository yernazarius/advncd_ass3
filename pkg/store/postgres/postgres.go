package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func DBconnection(username, password, localhost, port, dbname string) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", username, password, localhost, port, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}