package main

import (
	"fmt"
	"log"
	"my-project/beans"
	"net/http"

	"my-project/router"

	"github.com/gorilla/handlers"
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
	beans.Db.AutoMigrate(&beans.Student{})
	beans.Db.AutoMigrate(&beans.Teacher{})
	beans.Db.AutoMigrate(&beans.StudentAttendance{})
	log.Fatal(http.ListenAndServe(":8000",handlers.CORS( 
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	)(router.Router)))

}