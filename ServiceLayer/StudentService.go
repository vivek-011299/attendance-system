package ServiceLayer

import (
	"fmt"
	"my-project/RepoLayer"
	"my-project/beans"
)

func StudentPunchin(student_obj beans.StudentAttendance) string{
	stu_table := RepoLayer.SearchStudent(student_obj.StudentId)
	if stu_table == (beans.Student{}){
		return "No student found with that id"
	} else{
		count_obj := RepoLayer.Count_of_records(student_obj.StudentId)
		fmt.Println("count obj is :", count_obj)
		fmt.Println("length of count obj is ", len(count_obj))
		var flag bool = true
		for i := 0; i < len(count_obj); i++ {
			if count_obj[i].PunchOut == "0"{
				flag=false
			}
		}
		if flag{
			RepoLayer.InsertPunchin_time(student_obj)
			return "Inserted"
		}
		return "Error in inserting"
	}
}


func StudentPunchout(student_obj beans.StudentAttendance) string{
	stu_table := RepoLayer.SearchStudent(student_obj.StudentId)
	if stu_table == (beans.Student{}){
		return "No student found with that id"
	}else{
		count_obj := RepoLayer.Count_of_records(student_obj.StudentId)
		fmt.Println("count obj is :", count_obj)
		fmt.Println("length of count obj is ", len(count_obj))
		var flag bool = true
		for i := 0; i < len(count_obj); i++ {
			if count_obj[i].PunchOut=="0"{
				flag=false
			}
		}
		if flag{
			RepoLayer.InsertPunchOut(student_obj)
			return "Inserted"
		}else{
			return "Error in inserting punchout time"
		}
	}
}

func SearchStudent(stu_id int) beans.Student{
	student_object := RepoLayer.SearchStudent(stu_id)
	return student_object
}