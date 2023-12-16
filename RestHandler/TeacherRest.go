package RestHandler

import (
	"encoding/json"
	"fmt"
	"log"
	"my-project/ServiceLayer"
	"my-project/beans"
	"net/http"
	"net/url"
	"strconv"
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
		json.NewEncoder(w).Encode(student_details)
	}
}

func GetStudentAttendance(w http.ResponseWriter, r *http.Request){
	u, err := url.Parse(r.URL.RequestURI())
	if err!=nil{
		log.Fatal(err)
	}
	params, _ := url.ParseQuery(u.RawQuery)
	stu_id_string := params.Get("id")
	stu_id, _ := strconv.Atoi(stu_id_string)
	student_attendance := ServiceLayer.GetStudentAttendance(stu_id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student_attendance)
}
func GetStudentbyID(w http.ResponseWriter, r *http.Request){
	u, err := url.Parse(r.URL.RequestURI())
	if err!=nil{
		log.Fatal(err)
	}
	params, _ := url.ParseQuery(u.RawQuery)
	student_id,_ := strconv.Atoi(params.Get("id"))
	student_details := ServiceLayer.GetStudentbyID(student_id)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(student_details)
}

func CreateStudent(w http.ResponseWriter, r *http.Request){
	var student_obj beans.Student
	err := json.NewDecoder(r.Body).Decode(&student_obj)
	if err!=nil{
		log.Fatal("err in create student function", err)
	}
	message := ServiceLayer.CreateStudent(student_obj)
	fmt.Println("message from rest handler", message)
}