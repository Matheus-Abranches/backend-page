package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// conecta com o db
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error opening connection to the database: ", err)
		return nil, err
	}

	// verificar se deu bom
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database: ", err)
		return nil, err
	}

	fmt.Println("Connected to " + dbname)
	return db, nil
}
