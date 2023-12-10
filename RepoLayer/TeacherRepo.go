package RepoLayer

import (
	"fmt"
	"log"
	"my-project/beans"
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