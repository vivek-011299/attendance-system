package RepoLayer

import (
	"fmt"
	"my-project/beans"
)

func PrincipalGetAllStudents() []beans.Student{
		var stu_obj []beans.Student
		rows, err := beans.Db.DB().Query("SELECT * from students")
		if err!=nil{
			fmt.Println("error in query",err)
		}
		for rows.Next(){
			var student beans.Student
			err := rows.Scan(&student.Id, &student.Name, &student.Class, &student.Age,&student.Phone)
			if err!=nil{
				fmt.Println("err in scanning ", err)
			}
			stu_obj = append(stu_obj, student)
		}
		return stu_obj
}

func PrincipalGetStudentbyID(student_id string) beans.Student{
	var stu_obj beans.Student
	row := beans.Db.DB().QueryRow("Select * from students where id ="+student_id)
	err := row.Scan(&stu_obj.Id,&stu_obj.Name,&stu_obj.Class,&stu_obj.Age,&stu_obj.Phone)
	if err!=nil{
		fmt.Println("err in scanning")
		return beans.Student{}
	}
	return stu_obj
}