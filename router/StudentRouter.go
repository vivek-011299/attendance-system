package router

import (
	"my-project/RestHandler"

	"github.com/gorilla/mux"
)

func StudentRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/student", RestHandler.StudentPage)
	router.HandleFunc("/punchin", RestHandler.StudentPunchin).Methods("POST")
	router.HandleFunc("/punchout", RestHandler.StudentPunchout).Methods("POST")

}