package router

import (
	"my-project/RestHandler"

	"github.com/gorilla/mux"
)

func TeacherRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/get_all_students", RestHandler.GetAllStudents).Methods("GET")
	router.HandleFunc("/get_student/{id}", RestHandler.GetStudentbyID).Methods("GET")
	router.HandleFunc("/get_student_attendance/{id}", RestHandler.GetStudentAttendance).Methods("GET")
	router.HandleFunc("/create_student", RestHandler.CreateStudent).Methods("POST")

}