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

func StudentPunchin(w http.ResponseWriter, r *http.Request){
	var stu_punchin_obj beans.StudentAttendance
	err := json.NewDecoder(r.Body).Decode(&stu_punchin_obj)
	if err!=nil{
		fmt.Println("err in decoding obj", err)
	}
	var msg_obj beans.MessageBox
	msg_obj.Message = ServiceLayer.StudentPunchin(stu_punchin_obj)
	json.NewEncoder(w).Encode(msg_obj)
}

func StudentPunchout(w http.ResponseWriter, r *http.Request){
	var stu_punchin_obj beans.StudentAttendance
	err := json.NewDecoder(r.Body).Decode(&stu_punchin_obj)
	if err!=nil{
		fmt.Println("err in decoding obj during punchout", err)
	}
	var msg_obj beans.MessageBox
	msg_obj.Message = ServiceLayer.StudentPunchout(stu_punchin_obj)
	json.NewEncoder(w).Encode(msg_obj)
}

func SearchStudent(w http.ResponseWriter, r *http.Request){
	u, err := url.Parse(r.URL.RequestURI())
	fmt.Println(u)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(u.RawQuery)
	params, _ := url.ParseQuery(u.RawQuery)
	id,_ := strconv.Atoi(params.Get("id"))
	fmt.Println("id is ", id)
	student_search := ServiceLayer.SearchStudent(id)
	json.NewEncoder(w).Encode(student_search)
}
