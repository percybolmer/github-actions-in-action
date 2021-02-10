package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"testing"

	_ "github.com/lib/pq"
)

func TestThis(t *testing.T) {

}

func TestThis2(t *testing.T) {

}

// a Test that tries to connect to a PostgreSQL database
func TestConnect(t *testing.T) {
	// First get some ENV variabeles
	host := os.Getenv("host")
	user := os.Getenv("username")
	portStr := os.Getenv("port")
	password := os.Getenv("password")
	dbname := os.Getenv("database")

	// convert port to string
	port, err := strconv.Atoi(portStr)
	if err != nil {
		t.Fatal(err)
	}

	psqlConnection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open the DB connection
	db, err := sql.Open("postgres", psqlConnection)
	if err != nil {
		t.Fatal(err)
	}
	// Make sure connection really works
	err = db.Ping()
	if err != nil {
		t.Fatal(err)
	}
}
