package models

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const(
	USER = "postgres"
	PASS = "Testing1999"
	DBNAME = "test"
)
func Connect() *sql.DB{
	URL :=fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", USER, PASS, DBNAME)
	db, err := sql.Open("postgres", URL)
	if err!= nil {
		log.Fatal(err)
	}
	return db
}
func TestConnection()  {
	con := Connect()
	defer con.Close()
	err := con.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection is successful")
}
