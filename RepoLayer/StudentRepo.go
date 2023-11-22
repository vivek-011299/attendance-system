package RepoLayer

import (
	"fmt"
	"my-project/beans"
)


func InsertPunchin_time(student_obj beans.StudentAttendance){
	beans.Db.Exec("INSERT INTO student_attendance (roll, punchin) VALUES (?, ?)",student_obj.StudentId, student_obj.PunchIn )
}


func Get_recent_attendance_record(student_id int) beans.StudentAttendance{
	dayRecord := beans.StudentAttendance{}
	query := fmt.Sprintf("SELECT * FROM student_attendance order by punchin desc limit 1 where roll = %d", student_id)
	beans.Db.Raw(query).Scan(&dayRecord)
	return dayRecord
}

func Count_of_records(student_id int) beans.StudentAttendance{
	count_obj := beans.StudentAttendance{}
	beans.Db.Exec("SELECT * FROM student_attendance WHERE roll = ?",student_id).Scan(&count_obj)
	return count_obj
}
