package RepoLayer

import (
	"fmt"
	"log"
	"my-project/beans"
	"strconv"
)

	
func InsertPunchin_time(student_obj beans.StudentAttendance){
	query := fmt.Sprintf("INSERT INTO student_attendances VALUES (%d, '%s', '%s')",student_obj.StudentId, student_obj.PunchIn, student_obj.PunchOut)
	beans.Db.Exec(query)
}


func Get_recent_attendance_record(student_id int) beans.StudentAttendance{
	dayRecord := beans.StudentAttendance{}
	query := fmt.Sprintf("SELECT * FROM student_attendances where student_id = %d order by punch_in desc limit 1 ", student_id)
	beans.Db.Raw(query).Scan(&dayRecord)
	return dayRecord
}

func Count_of_records(student_id int) []beans.StudentAttendance{
	var stu_count_obj []beans.StudentAttendance
	rows, err := beans.Db.DB().Query("SELECT student_id, punch_in, punch_out FROM student_attendances WHERE student_id = "+ strconv.Itoa(student_id))
	if err!=nil{
		log.Fatal("Error is here", err)
	}
	for rows.Next() {
        var stu_count beans.StudentAttendance
		if err := rows.Scan(&stu_count.StudentId, &stu_count.PunchIn, &stu_count.PunchOut); err != nil {
            fmt.Println("error in scanning", err)
        }
        stu_count_obj = append(stu_count_obj, stu_count)
    }
    fmt.Println("stu_count is ",stu_count_obj)
	return stu_count_obj
}


func InsertPunchOut(student_obj beans.StudentAttendance) {
	query := fmt.Sprintf("UPDATE student_attendances SET punch_out = '%s' WHERE student_id = %d and punch_out='0'",student_obj.PunchOut, student_obj.StudentId)
	beans.Db.Exec(query)
}


func SearchStudent(student_id int) beans.Student{
	var student_table beans.Student
	row := beans.Db.DB().QueryRow("SELECT * from students where id = "+ strconv.Itoa(student_id))
    err := row.Scan(&student_table.Id, &student_table.Name,&student_table.Class, &student_table.Age,&student_table.Phone)
	if err!=nil{
		return beans.Student{}
	}
	return student_table
}