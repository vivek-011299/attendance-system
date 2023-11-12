package main

import (
	"database/sql"
	"fmt"
	"os"
)

func OpenDatabase(db *sql.DB) error{
	db_name := os.Getenv("DB_NAME")
	username := os.Getenv("USER")
	driver := os.Getenv("DRIVER_NAME")
	db, err := sql.Open(driver, "user="+username+"dbname="+db_name+"sslmode=disable")
	fmt.Println("db is ",db)
	if err!=nil{
		return (err)
	}
	return nil
}

func main() {
	//router := mux.NewRouter()
	var db *sql.DB
	err := OpenDatabase(db)
	if err!=nil{
		fmt.Printf("error connecting %v",err)
	}

	fmt.Println("Starting the server at port 8000")
}