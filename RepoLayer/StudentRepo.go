package RepoLayer

import (
	"fmt"
	"log"
	"my-project/beans"
	"strconv"
)

	
func InsertPunchin_time(student_obj beans.StudentAttendance){
	beans.Db.Exec("INSERT INTO student_attendance (roll, punchin) VALUES (?, ?)",student_obj.StudentId, student_obj.PunchIn )
}


func Get_recent_attendance_record(student_id int) beans.StudentAttendance{
	dayRecord := beans.StudentAttendance{}
	query := fmt.Sprintf("SELECT * FROM student_attendance where roll = %d order by punchin desc limit 1 ", student_id)
	beans.Db.Raw(query).Scan(&dayRecord)
	return dayRecord
}

func Count_of_records(student_id int) []beans.StudentAttendance{
	var stu_count_obj []beans.StudentAttendance
	rows, err := beans.Db.DB().Query("SELECT roll, punchin FROM student_attendance WHERE roll = "+ strconv.Itoa(student_id))
	if err!=nil{
		log.Fatal("Error is ", err)
	}
	for rows.Next() {
        var stu_count beans.StudentAttendance
		if err := rows.Scan(&stu_count.StudentId, &stu_count.PunchIn); err != nil {
            fmt.Println("error in scanning", err)
        }
        stu_count_obj = append(stu_count_obj, stu_count)
    }
    fmt.Println("stu_count is ",stu_count_obj)
	return stu_count_obj
}


func InsertPunchOut(student_obj beans.StudentAttendance) {
	beans.Db.Exec("UPDATE student_attendance SET punchout = ? WHERE roll = ?",student_obj.PunchOut, student_obj.StudentId)
}


func SearchStudent(student_id int) []beans.Student{
	var student_table []beans.Student
	rows, err := beans.Db.DB().Query("SELECT * from students where id = "+ strconv.Itoa(student_id))
	if err!=nil{
		log.Fatal("Error in searching student", err)
	}
	for rows.Next() {
        var stu_count beans.Student
		if err := rows.Scan(&stu_count.Id, &stu_count.Name ,&stu_count.Class, &stu_count.Age, &stu_count.Phone); err != nil {
            fmt.Println("error in scanning", err)
        }
        student_table = append(student_table, stu_count)
    }
	return student_table
}