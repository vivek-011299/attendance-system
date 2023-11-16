package router

import (
	"my-project/RestHandler"

	"github.com/gorilla/mux"
)

func PrincipalRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/get_all_students", RestHandler.PrincipalGetAllStudents).Methods("GET")
	router.HandleFunc("/get_student/{id}", RestHandler.PrincipalGetStudentbyID).Methods("GET")
	router.HandleFunc("/get_all_teachers", RestHandler.GetAllTeachers).Methods("GET")
	router.HandleFunc("/get_teacher/{id}", RestHandler.GetTeacherbyID).Methods("GET")
	router.HandleFunc("/create_teacher", RestHandler.CreateTeacher).Methods("POST")

	
}