package RepoLayer

import (
	"fmt"
	"log"
	"my-project/beans"
	"strconv"
)

func GetAllStudents() []beans.Student {
	var students []beans.Student
	rows, err:= beans.Db.DB().Query("select * from students")
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

func GetStudentbyID(student_id int) beans.Student{
	var student_details beans.Student
	row := beans.Db.DB().QueryRow("Select * from students where id = "+strconv.Itoa(student_id))
	err := row.Scan(&student_details.Id, &student_details.Name,&student_details.Class,&student_details.Age,&student_details.Phone)
	if err!=nil{
		return beans.Student{}
	}
	return student_details
}


func InsertStudent(student_obj beans.Student)string{
	name := student_obj.Name
	class := strconv.Itoa(student_obj.Class)
	age := strconv.Itoa(student_obj.Age)
	phone := student_obj.Phone
	q := fmt.Sprintf("INSERT INTO students(name,class,age,phone) VALUES('%s',%s,%s,'%s')",name,class,age,phone)
	beans.Db.Exec(q)
	return "inserted"
}

func Delete_Student(stu_id int){
	student_id := strconv.Itoa(stu_id)
	beans.Db.Exec("DELETE from students where id="+student_id)
	fmt.Println("deleted record with id "+student_id)
}