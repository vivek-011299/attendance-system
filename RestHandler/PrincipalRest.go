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
	teacher_obj := ServiceLayer.GetAllTeachers()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teacher_obj)
}

func GetTeacherbyID(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.RequestURI())
	if err!=nil{
		log.Fatal(err)
	}
	params, _ := url.ParseQuery(u.RawQuery)
	t_id := params.Get("id")
	teacher_details := ServiceLayer.GetTeacherbyID(t_id)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(teacher_details)
}

func CreateTeacher(w http.ResponseWriter, r *http.Request) {
	var teacher_obj beans.Teacher
	err := json.NewDecoder(r.Body).Decode(&teacher_obj)
	if err!=nil{
		fmt.Println("err in decoding", err);
	}
	message := ServiceLayer.CreateTeacher(teacher_obj)
	fmt.Println(message)
	json.NewEncoder(w).Encode(teacher_obj)
}

func Delete_Teacher(w http.ResponseWriter, r *http.Request){
	u, err := url.Parse(r.URL.RequestURI())
	if err!=nil{
		fmt.Println("err in parsing", err)
	}
	params,_ := url.ParseQuery(u.RawQuery)
	t_id,_ := strconv.Atoi(params.Get("id"))
	ServiceLayer.Delete_Teacher(t_id) 
}