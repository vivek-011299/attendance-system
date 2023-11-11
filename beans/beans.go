package beans

import "time"

type Student struct {
	Id    int    `json:"roll"`
	Name  string `json:"name"`
	Class int    `json:"class"`
	Age   int    `json:"age"`
	Phone string `json:"phone"`
}

type Teacher struct {
	Id          int    `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	PhoneNumber string `json:"phone number"`
}

type StudentAttendance struct {
	StudentId int `json:"roll"`
	PunchIn  time.Time `json:"punchin"`
	PunchOut time.Time `json:"punchout"`
}