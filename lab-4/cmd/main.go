package main

import (
	"database/sql"
	"lab-4/internal/api/controller"
	"lab-4/internal/api/router"
	"lab-4/internal/infra/database"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		Addr:   os.Getenv("MYSQL_ADDRESS"),
		DBName: os.Getenv("MYSQL_DATABASE"),
		User:   os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		Net:    "tcp",
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	repository := database.NewBookRepository(db)
	controller := controller.NewBookController(repository)
	router.StartBookRouter(":8080", controller)
}
