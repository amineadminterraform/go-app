package main

import (
	"context"
	"log"

	"github.com/amineadminterraform/go-app/api"
	db "github.com/amineadminterraform/go-app/db/sqlc"
	"github.com/amineadminterraform/go-app/utils"
	"github.com/jackc/pgx/v5"
)



func main(){
	
	config , err := utils.LoadConfig("./")
	if err != nil{
		log.Fatal("cannot load config :", err)
	}
	testDB,err:= pgx.Connect(context.Background(),config.DBSource)
	if err!=nil{
		log.Fatal("You cannot connect to db :", err)
	}
	defer testDB.Close(context.Background())
	store := db.NewStore(testDB)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if (err != nil){
		log.Fatal("cannot start server :", err)
	}	
}