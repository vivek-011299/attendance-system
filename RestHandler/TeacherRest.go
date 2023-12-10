package RestHandler

import (
	"encoding/json"
	"fmt"
	"log"
	"my-project/ServiceLayer"
	"net/http"
	"net/url"
)

func GetAllStudents(w http.ResponseWriter, r *http.Request){
	student_details := ServiceLayer.GetAllStudents()
	if len(student_details)==0{
		fmt.Println("No records found")
	}else{
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(student_details)
	}
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