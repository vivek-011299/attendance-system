package ServiceLayer

import (
	"my-project/RepoLayer"
	"my-project/beans"
)

func GetAllStudents() []beans.Student{
	Student_details:= RepoLayer.GetAllStudents()
	return Student_details
}

func GetStudentAttendance(student_id int) []beans.StudentAttendance{
	student_attendance_details := RepoLayer.GetStudentAttendance(student_id)
	return student_attendance_details
}