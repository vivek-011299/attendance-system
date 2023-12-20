package ServiceLayer

import (
	"fmt"
	"my-project/RepoLayer"
	"my-project/beans"
)

func StudentPunchin(student_obj beans.StudentAttendance){
	count_obj := RepoLayer.Count_of_records(student_obj.StudentId)
	fmt.Println("count obj is :", count_obj)
	fmt.Println("length of count obj is ", len(count_obj))
	var flag bool = true
	for i := 0; i < len(count_obj); i++ {
		if count_obj[i].PunchOut.IsZero(){
			flag=false
		}
	}
	if len(count_obj)==0{
		fmt.Println("No student present with student id ",student_obj.StudentId)
	}else{ 
		if flag{
			fmt.Println("inserting time")
			RepoLayer.InsertPunchin_time(student_obj)
		}else{
			fmt.Println("Please punchout for your previous day")
		}
	}
}


func StudentPunchout(student_obj beans.StudentAttendance){
	count_obj := RepoLayer.Count_of_records(student_obj.StudentId)
	fmt.Println("count obj is :", count_obj)
	fmt.Println("length of count obj is ", len(count_obj))
	var flag bool = true
	for i := 0; i < len(count_obj); i++ {
		if count_obj[i].PunchOut.IsZero(){
			flag=false
		}
	}
	if len(count_obj)==0{
		fmt.Println("No student present with student id ",student_obj.StudentId)
	}else{
		if flag{
			fmt.Println("Please punchin before punching out.")
		}else{
			RepoLayer.InsertPunchOut(student_obj)
		}
	}
}

func SearchStudent(stu_id int) beans.Student{
	student_object := RepoLayer.SearchStudent(stu_id)
	return student_object
}