package RepoLayer

import (
	"fmt"
	"my-project/beans"
)


func InsertPunchin_time(student_obj beans.StudentAttendance){
	query := fmt.Sprintf("INSERT INTO student_attendance (roll, punchin) VALUES (%d, %v)", student_obj.StudentId, student_obj.PunchIn)
	result := beans.Db.Raw(query)
	fmt.Println("result is : ",result.Value)
}


func Get_recent_attendance_record(student_id int) beans.StudentAttendance{
	dayRecord := beans.StudentAttendance{}
	query := fmt.Sprintf("SELECT * FROM student_attendance order by punchin desc limit 1 where roll = %d", student_id)
	beans.Db.Raw(query).Scan(&dayRecord)
	return dayRecord
}

func Count_of_records(student_id int) beans.StudentAttendance{
	count := beans.StudentAttendance{}
	query := fmt.Sprintf("SELECT * FROM student_attendance WHERE roll = %d",student_id)
	beans.Db.Raw(query).Scan(&count)
	//beans.Db.Table("student_attendance").Select("*").Where("roll = ", student_id).Scan(&count)
	return count
}