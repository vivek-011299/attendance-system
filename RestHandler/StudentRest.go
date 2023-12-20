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
	"time"
)

func StudentPunchin(w http.ResponseWriter, r *http.Request){
	u, err:= url.Parse(r.URL.RequestURI())
	if err!=nil{
		log.Fatal(err)
	}
	params,_ := url.ParseQuery(u.RawQuery)
	fmt.Println(params["id"])
	stu_id,_ := strconv.Atoi(params.Get("id"))
	student_punchin_obj := beans.StudentAttendance{
		StudentId: stu_id,
		PunchIn: time.Now(),
	}
	ServiceLayer.StudentPunchin(student_punchin_obj)
}

func StudentPunchout(w http.ResponseWriter, r *http.Request){
		u,err := url.Parse(r.URL.RequestURI())
		if err!=nil{
			log.Fatal("error in student punchout", err)
		}
		params, _ := url.ParseQuery(u.RawQuery)
		fmt.Println(params)
		stu_id,_ := strconv.Atoi(params.Get("id"))
		student_punchout_obj := beans.StudentAttendance{
			StudentId: stu_id,
			PunchOut: time.Now(),
		}
		ServiceLayer.StudentPunchout(student_punchout_obj)
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
