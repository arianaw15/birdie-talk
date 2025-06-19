package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/arianaw15/birdie-talk/cmd/api"
	"github.com/arianaw15/birdie-talk/config"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := config.Envs.DBUser + ":" + config.Envs.DBPassword + "@tcp(127.0.0.1:3306)/birdie_talk?charset=utf8"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to database")

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
