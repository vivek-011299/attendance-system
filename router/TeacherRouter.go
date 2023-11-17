package router

import (
	"my-project/RestHandler"
)

func TeacherRouter() {
	Router.HandleFunc("/teacher/get_all_students", RestHandler.GetAllStudents).Methods("GET")
	Router.HandleFunc("/teacher/get_student", RestHandler.GetStudentbyID).Methods("GET")
	Router.HandleFunc("/teacher/get_student_attendance", RestHandler.GetStudentAttendance).Methods("GET")
	Router.HandleFunc("/teacher/create_student", RestHandler.CreateStudent).Methods("POST")
}