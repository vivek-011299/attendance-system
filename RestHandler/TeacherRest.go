package RestHandler

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func TeacherPage(w http.ResponseWriter, r *http.Request){
	fmt.Println("Welcome to teacher page")
}

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got all students")
}

func GetStudentAttendance(w http.ResponseWriter, r *http.Request){
	u, err := url.Parse(r.URL.RequestURI())
	if err!=nil{
		log.Fatal(err)
	}
	params, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(params)
}
func GetStudentbyID(w http.ResponseWriter, r *http.Request){
	u, err := url.Parse(r.URL.RequestURI())
	if err!=nil{
		log.Fatal(err)
	}
	params, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(params)
}

func CreateStudent(w http.ResponseWriter, r *http.Request){
	fmt.Println("crete a student")
}