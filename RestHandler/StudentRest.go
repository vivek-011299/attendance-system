package RestHandler

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func StudentPage(w http.ResponseWriter, r *http.Request){
	fmt.Println("Here comes student page")
}

func StudentPunchin(w http.ResponseWriter, r *http.Request){

	fmt.Println("Student punched in")
}

func StudentPunchout(w http.ResponseWriter, r *http.Request){
	fmt.Println("student punched out")
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
