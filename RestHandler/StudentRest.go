package RestHandler

import (
	"fmt"
	"log"
	"my-project/ServiceLayer"
	"my-project/beans"
	"net/http"
	"net/url"
	"time"
)

func StudentPage(w http.ResponseWriter, r *http.Request){
	fmt.Println("Here comes student page")
}


func StudentPunchin(w http.ResponseWriter, r *http.Request){
	u, err:= url.Parse(r.URL.RequestURI())
	if err!=nil{
		log.Fatal(err)
	}
	params,_ := url.ParseQuery(u.RawQuery)
	fmt.Println(params["id"])
	stu_id := params.Get("id")
	_, ok := beans.Pipout[stu_id]
	if ok {
		fmt.Println("You need to punchout first before punching in")
	}else{
		beans.Pipout[stu_id] = time.Now()
	}
}

func StudentPunchout(w http.ResponseWriter, r *http.Request){
		u,err := url.Parse(r.URL.RequestURI())
		if err!=nil{
			log.Fatal("error in student punchout", err)
		}
		params, _ := url.ParseQuery(u.RawQuery)
		fmt.Println(params)
		stu_id := params.Get("id")
		_, ok := beans.Pipout[stu_id]
		if ok{
			student_obj := beans.StudentAttendance{
				StudentId: stu_id,
				PunchIn: beans.Pipout[stu_id],
				PunchOut: time.Now(),
			}
			delete(beans.Pipout, stu_id)
			ServiceLayer.StudentPunchout(student_obj)
		}else{
			fmt.Println("You need to punchin first")
		}
}

func SearchStudent(w http.ResponseWriter, r *http.Request){
	u, err := url.Parse(r.URL.RequestURI())
	fmt.Println(u)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(u.RawQuery)
	params, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(params)
}
