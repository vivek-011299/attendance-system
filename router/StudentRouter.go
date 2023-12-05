package router

import (
	"my-project/RestHandler"
)

func StudentRouter() {
	Router.HandleFunc("/student/search", RestHandler.SearchStudent)
	Router.HandleFunc("/student/punchin", RestHandler.StudentPunchin)
	Router.HandleFunc("/student/punchout", RestHandler.StudentPunchout)
}