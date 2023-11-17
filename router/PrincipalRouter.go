package router

import (
	"my-project/RestHandler"
)

func PrincipalRouter() {
	Router.HandleFunc("/principal/get_all_students", RestHandler.PrincipalGetAllStudents).Methods("GET")
	Router.HandleFunc("/principal/get_student/{id}", RestHandler.PrincipalGetStudentbyID).Methods("GET")
	Router.HandleFunc("/principal/get_all_teachers", RestHandler.GetAllTeachers).Methods("GET")
	Router.HandleFunc("/principal/get_teacher/{id}", RestHandler.GetTeacherbyID).Methods("GET")
	Router.HandleFunc("/principal/create_teacher", RestHandler.CreateTeacher).Methods("POST")

	
}