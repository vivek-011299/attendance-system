package ServiceLayer

import (
	"fmt"
	"my-project/RepoLayer"
	"my-project/beans"
)

func StudentPunchin(student_obj beans.StudentAttendance){
	latestDayRecord := beans.StudentAttendance{}
	latestDayRecord = RepoLayer.Get_recent_attendance_record(student_obj.StudentId)
	if latestDayRecord.PunchOut.IsZero(){
		fmt.Println("You need to punchout first for the previous day.")
	}else{
		RepoLayer.InsertPunchin_time(student_obj.StudentId)
	}
}


func StudentPunchout(student_obj beans.StudentAttendance){

}