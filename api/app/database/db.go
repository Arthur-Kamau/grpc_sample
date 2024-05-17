package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func DbInstance(dbname string) *sql.DB {

	fmt.Printf("Username %s Password %s Host %s Port %s db Name %s  \n", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), dbname)
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		dbname)

	fmt.Printf("\n Connection string %s", connString)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Printf("error openning db %s ", err.Error())
		return nil
	}
	
  
	return db

}
