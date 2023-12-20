package router

import (
	"my-project/RestHandler"
)

func StudentRouter() {
	Router.HandleFunc("/student/search", RestHandler.SearchStudent).Methods("GET")
	Router.HandleFunc("/student/punchin", RestHandler.StudentPunchin).Methods("POST")
	Router.HandleFunc("/student/punchout", RestHandler.StudentPunchout).Methods("POST")
}