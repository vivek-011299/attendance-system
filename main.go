package main

import (
	"fmt"
	"log"
	"my-project/beans"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var db *gorm.DB
var err error

func main() {
	//loading env variables
	dialect := os.Getenv("DRIVER")
	//db_name := os.Getenv("DBNAME")
	//db_port := os.Getenv("DBPORT")
	//user := os.Getenv("USER")
	//password := os.Getenv("PASSWORD")
	//host:= os.Getenv("HOST")
	dbURI := "host=localhost user=postgres dbname=attendance sslmode=disable password=postgres port=5432"

	//opening db
	db, err = gorm.Open(dialect,dbURI)
	if err!=nil{
		log.Fatal(err)
	}else{
		fmt.Println("Starting the server at port 5432")
	}
	
	defer db.Close()

	db.AutoMigrate(&beans.Student{})
	db.AutoMigrate(&beans.Teacher{})
	db.AutoMigrate(&beans.StudentAttendance{})
}