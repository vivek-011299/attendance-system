package ServiceLayer

import (
	"my-project/RepoLayer"
	"my-project/beans"
)

func GetAllStudents() []beans.Student{
	Student_details:= RepoLayer.GetAllStudents()
	return Student_details
}