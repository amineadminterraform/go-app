package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/platform_transactions?sslmode=disable" 
)

	
func TestMain(m *testing.M){

	conn,err:= pgx.Connect(context.Background(),dbSource)
	if err!=nil{
		log.Fatal("You cannot connect to db :", err)
	}
	defer conn.Close(context.Background())
	testQueries= New(conn)
	os.Exit(m.Run())
}

