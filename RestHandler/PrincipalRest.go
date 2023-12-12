package RestHandler

import (
	"encoding/json"
	"fmt"
	"log"
	"my-project/ServiceLayer"
	"net/http"
	"net/url"
)


func PrincipalGetAllStudents(w http.ResponseWriter, r *http.Request) {
	students_obj := ServiceLayer.PrincipalGetAllStudents()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students_obj)
}

func PrincipalGetStudentbyID(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.RequestURI())
	if err!=nil{
		log.Fatal(err)
	}
	params, _ := url.ParseQuery(u.RawQuery)
	student_id := params.Get("id")
	student_obj := ServiceLayer.PrincipalGetStudentbyID(student_id)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(student_obj)
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