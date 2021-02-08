package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	connectionString := "gripuser:grip2020user@tcp(127.0.0.1:3306)grip"
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var name string
}