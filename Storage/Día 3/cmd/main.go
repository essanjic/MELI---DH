package main

import (
	"database/sql"
	//go get "github.com/go-sql-driver/mysql" esto se ejecuta en la terminal
	"github.com/go-sql-driver/mysql"
)

func main() {
	databaseConfig := mysql.Config{
		User:     "root",
		Password: "SuperSecretPassword123",
		Addr:     "127.0.0.1:3306",
		DBName:   "my_database",
	}
	database, err := sql.Open("mysql", databaseConfig.FormatDSN())
	if err != nil {
		panic(err)
	}
	defer database.Close()

	if err := database.Ping(); err != nil {
		panic(err)
	}

	defer database.Close()
}
