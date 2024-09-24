package main

import (
	"context"
	"log"

	"github.com/amineadminterraform/go-app/api"
	db "github.com/amineadminterraform/go-app/db/sqlc"
	_ "github.com/amineadminterraform/go-app/docs"
	"github.com/amineadminterraform/go-app/utils"
	"github.com/jackc/pgx/v5"
)

// @title PLATFORM OPERATIONS API
// @description A PLATFORM OPERATIONS API in golang using Gin framework
// @version 1.0
// @host localhost:8000
// @BasePath /
func main() {

	config, err := utils.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load config :", err)
	}
	testDB, err := pgx.Connect(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("You cannot connect to db :", err)
	}
	defer testDB.Close(context.Background())
	store := db.NewStore(testDB)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server :", err)
	}
	log.Println("Server Starting at localhost:8000")
	log.Println("Swagger at  localhost:8000/docs/index.html")

}
