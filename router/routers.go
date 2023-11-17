package router

import (
	"fmt"
	"my-project/RestHandler"

	"github.com/gorilla/mux"
)

var Router *mux.Router

func InitRouters(){
	fmt.Println("Inside InitRouters")
	Router = mux.NewRouter()
	Router.HandleFunc("/student", RestHandler.StudentPage)
	Router.HandleFunc("/teacher", RestHandler.TeacherPage)
	Router.HandleFunc("/principal", RestHandler.PrincipalPage)
	StudentRouter()
	TeacherRouter()
	PrincipalRouter()
}