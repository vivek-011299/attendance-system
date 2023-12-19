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

func GetAllTeachers() []beans.Teacher{
	teacher_obj := RepoLayer.GetAllTeachers()
	return teacher_obj
}

func GetTeacherbyID(t_id string) beans.Teacher{
	t_obj := RepoLayer.GetTeacherbyID(t_id)
	return t_obj
}

func CreateTeacher(teacher_obj beans.Teacher) string{
	if teacher_obj.Id != 0{
		return "dont pass id"
	}else{
		RepoLayer.InsertTeacher(teacher_obj)
	}
	return "inserted"
}

func Delete_Teacher(teacher_id int){
	RepoLayer.Delete_Teacher(teacher_id)
}