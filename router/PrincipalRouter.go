package router

import (
	"my-project/RestHandler"
)

func PrincipalRouter() {
	Router.HandleFunc("/principal/get_all_students", RestHandler.PrincipalGetAllStudents).Methods("GET")
	Router.HandleFunc("/principal/get_student", RestHandler.PrincipalGetStudentbyID).Methods("GET")
	Router.HandleFunc("/principal/get_all_teachers", RestHandler.GetAllTeachers).Methods("GET")
	Router.HandleFunc("/principal/get_teacher", RestHandler.GetTeacherbyID).Methods("GET")
	Router.HandleFunc("/principal/create_teacher", RestHandler.CreateTeacher).Methods("POST")	
	Router.HandleFunc("/principal/delete_teacher", RestHandler.Delete_Teacher).Methods("DELETE")
}