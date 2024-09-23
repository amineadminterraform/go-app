package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/amineadminterraform/go-app/utils"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {

	Config, err := utils.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config :", err)
	}
	testDB, err := pgx.Connect(context.Background(), Config.DBSource)
	if err != nil {
		log.Fatal("You cannot connect to db :", err)
	}
	defer testDB.Close(context.Background())
	testQueries = New(testDB)
	os.Exit(m.Run())
}
