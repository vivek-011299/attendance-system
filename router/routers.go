package router

import (
	"fmt"

	"github.com/gorilla/mux"
)

var Router *mux.Router

func InitRouters(){
	fmt.Println("Inside InitRouters")
	Router = mux.NewRouter()
	StudentRouter()
	TeacherRouter()
	PrincipalRouter()
}