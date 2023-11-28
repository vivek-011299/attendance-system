package ServiceLayer

import (
	"fmt"
	"my-project/RepoLayer"
	"my-project/beans"
)

func StudentPunchin(student_obj beans.StudentAttendance){
	fmt.Println("in service")
	//count_obj := 
	RepoLayer.Count_of_records(student_obj.StudentId)
	//fmt.Println("count obj is :", count_obj)
	/*
	if count_obj.PunchIn.IsZero() && count_obj.PunchOut.IsZero(){
		fmt.Println("inserting time")
		RepoLayer.InsertPunchin_time(student_obj)
	}else{
		latestDayRecord := RepoLayer.Get_recent_attendance_record(student_obj.StudentId)
		fmt.Println("Latest dy record is :",latestDayRecord)
		if latestDayRecord.PunchOut.IsZero(){
			fmt.Println("You need to punchout first for the previous day.")
		}else{
			RepoLayer.InsertPunchin_time(student_obj)
		}
	}
	*/
}


func StudentPunchout(student_obj beans.StudentAttendance){

}
