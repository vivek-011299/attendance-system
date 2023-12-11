package RepoLayer

import (
	"fmt"
	"log"
	"my-project/beans"
	"strconv"
)

func GetAllStudents() []beans.Student {
	var students []beans.Student
	rows, err:= beans.Db.DB().Query("select * from Students")
	if err!=nil{
		log.Fatal(err)
	}
	for rows.Next(){
		var stu beans.Student
		err := rows.Scan(&stu.Id, &stu.Name, &stu.Class, &stu.Age, &stu.Phone)
		if err != nil {
            fmt.Println("error in scanning", err)
        }
        students = append(students, stu)
	}
	return students
}

func GetStudentAttendance(student_id int) []beans.StudentAttendance{
	var student_attendance_details []beans.StudentAttendance
	rows, err := beans.Db.DB().Query("Select * from student_attendance where roll = "+ strconv.Itoa(student_id))

	if err!=nil{
		log.Fatal("error in rows", err)
	}
	for rows.Next(){
		var student_attendance_obj beans.StudentAttendance
		err := rows.Scan(&student_attendance_obj.StudentId, &student_attendance_obj.PunchIn,&student_attendance_obj.PunchOut)
		if err!=nil{
			fmt.Println("error in scanning", err)
		}
		student_attendance_details = append(student_attendance_details, student_attendance_obj)
	}
	return student_attendance_details
}