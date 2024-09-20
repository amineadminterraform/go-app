package utils

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/exp/rand"
)
const alf = "abcdefghijklmnopqrstuvwxyz"
type RandomData struct {
		this   int64  `json:"this"`
		is string `json:"is"`
		fortesting  int64  `json:"fortesting"`
	}
func init() {
	rand.Seed(uint64(time.Now().UnixNano()))
}

//RANDOM GENERATES A RANDOM INTEGER BETWEEN min and max
func RandomInt(min,max int64) int64	{
	return min + rand.Int63n(max-min + 1 ) 
}

func RandomString(n int) string {
	var sb strings.Builder
	alf_len := len(alf)
	for i := 0; i < n; i++  {
		c := alf[rand.Intn(alf_len)]
		sb.WriteByte(c)
	}
	return sb.String();
}

func RandomGit() string {
	return "github.com/"+RandomString(5)+"/"+RandomString(5)
}
func RandomGitBranch() string {
	return RandomString(5)+"/"+RandomString(5)
}

func RandomName() string {
	return RandomString(10)
}				
func RandomDescription() pgtype.Text {
	return pgtype.Text{String: RandomString(30),  Valid:true}
}	
func RandomS3Path() string {
	return RandomString(5)+"/"+RandomString(5)
}

func RandomJSON() string {
	// Struct to hold random data
		data := RandomData{
		this:   RandomInt(1, 1000),
		is: RandomString(10),
		fortesting:  RandomInt(18, 80),
	}

	// Convert struct to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error generating JSON:", err)
		return ""
	}

	return string(jsonData)
}