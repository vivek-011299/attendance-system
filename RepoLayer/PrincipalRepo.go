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

func GetAllTeachers() []beans.Teacher{
	var teachers_final []beans.Teacher
	rows, err := beans.Db.DB().Query("Select * from teachers")
	if err!=nil{
		fmt.Println("err in query", err)
	}
	for rows.Next(){
		var teach beans.Teacher
		err := rows.Scan(&teach.Id, &teach.FirstName, &teach.LastName,&teach.PhoneNumber)
		if err!=nil{
			fmt.Println("err in scanning", err)
		}
		teachers_final = append(teachers_final, teach)
	}
	return teachers_final
}

func GetTeacherbyID(t_id string) beans.Teacher{
	var teacher_detail beans.Teacher
	row := beans.Db.DB().QueryRow("SELECT * from teachers where id = "+t_id)
	err := row.Scan(&teacher_detail.Id, &teacher_detail.FirstName,&teacher_detail.LastName,&teacher_detail.PhoneNumber)
	if err!=nil{
		fmt.Println("err in scanning ", err)
		return beans.Teacher{}
	}
	return teacher_detail
}

func InsertTeacher(t_obj beans.Teacher) {
	first_name := t_obj.FirstName
	last_name := t_obj.LastName
	phone := t_obj.PhoneNumber
	query := fmt.Sprintf("INSERT INTO teachers(first_name,last_name,phone_number) VALUES('%s','%s','%s')",first_name,last_name,phone)
	beans.Db.Exec(query)
}