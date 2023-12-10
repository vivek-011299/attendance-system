package RestHandler

import (
	"encoding/json"
	"fmt"
	"log"
	"my-project/ServiceLayer"
	"my-project/beans"
	"net/http"
	"net/url"
)

func GetAllStudents(w http.ResponseWriter, r *http.Request){
	student_details := ServiceLayer.GetAllStudents()
	w.Header().Set("Content-Type","application/json")
	if len(student_details)==0{
		data := beans.Student{
				Id:0,
				Name:"null",
				Class:0,
				Age:0,
				Phone:"null",	
		}
		json.NewEncoder(w).Encode(data)

	}else{
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