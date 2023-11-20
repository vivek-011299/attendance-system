package RepoLayer

import (
	"fmt"
	"my-project/beans"
)


func InsertPunchin_time(student_obj beans.StudentAttendance){
	query := fmt.Sprintf("INSERT INTO student_attendance (roll, punchin) VALUES (%d, %v)", student_obj.StudentId, student_obj.PunchIn)
	result := beans.Db.Exec(query, student_obj.StudentId, student_obj.PunchIn)
	rowsAffected := result.RowsAffected
	fmt.Println(rowsAffected)

}


func Get_recent_attendance_record(student_id int) beans.StudentAttendance{
	dayRecord := beans.StudentAttendance{}
	query := "SELECT * from student_attendance order by DESC limit 1"
	row := beans.Db.DB().QueryRow(query)
	if row == nil{
		fmt.Println("Error fetching the recent attendance record")
	}else{
		fmt.Println("query executed successfully.")
	}
	return dayRecord
}