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


func main() {
	dbURI := "host=localhost user=postgres dbname=attendance sslmode=disable password=postgres port=5432"

	//opening db
	var err error
	beans.Db, err = gorm.Open("postgres",dbURI)
	if err!=nil{
		log.Fatal(err)
	}else{
		fmt.Println("Starting the db at port 5432")
	}
	
	router.InitRouters()
	fmt.Println("Starting the server at 8000")
	defer beans.Db.Close()
	log.Fatal(http.ListenAndServe(":8000",router.Router))

	beans.Db.AutoMigrate(&beans.Student{})
	beans.Db.AutoMigrate(&beans.Teacher{})
	beans.Db.AutoMigrate(&beans.StudentAttendance{})
}