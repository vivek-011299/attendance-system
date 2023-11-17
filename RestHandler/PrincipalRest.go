package RestHandler

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func PrincipalPage(w http.ResponseWriter, r *http.Request){
	fmt.Println("welcome to principl page")
}

func PrincipalGetAllStudents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got all students")
}

func PrincipalGetStudentbyID(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.RequestURI())
	if err!=nil{
		log.Fatal(err)
	}
	params, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(params)
}

func GetAllTeachers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got all teachers")
}

func GetTeacherbyID(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.RequestURI())
	if err!=nil{
		log.Fatal(err)
	}
	params, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(params)
}

func CreateTeacher(w http.ResponseWriter, r *http.Request) {
	fmt.Println("created teacher")
}