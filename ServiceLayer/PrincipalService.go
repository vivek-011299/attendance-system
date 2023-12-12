package ServiceLayer

import (
	"my-project/RepoLayer"
	"my-project/beans"
)

func PrincipalGetAllStudents() []beans.Student{
	students_obj := RepoLayer.PrincipalGetAllStudents()
	return students_obj
}

func PrincipalGetStudentbyID(stu_id string) beans.Student{
	student_obj := RepoLayer.PrincipalGetStudentbyID(stu_id)
	return student_obj
}