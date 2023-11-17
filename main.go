package main

import (
	"fmt"
	"log"
	"my-project/beans"
	"net/http"

	"my-project/router"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var db *gorm.DB
var err error

func main() {
	dbURI := "host=localhost user=postgres dbname=attendance sslmode=disable password=postgres port=5432"

	//opening db
	db, err = gorm.Open("postgres",dbURI)
	if err!=nil{
		log.Fatal(err)
	}else{
		fmt.Println("Starting the server at port 5432")
	}
	
	router.InitRouters()
	defer db.Close()
	log.Fatal(http.ListenAndServe(":8000",router.Router))

	db.AutoMigrate(&beans.Student{})
	db.AutoMigrate(&beans.Teacher{})
	db.AutoMigrate(&beans.StudentAttendance{})
}