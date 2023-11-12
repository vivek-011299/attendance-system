package beans

import "time"

type Student struct {
	Id    int    `json:"roll" gorm:"primary_key; auto_increment"`
	Name  string `json:"name" gorm:"type:varchar(40); not null"`
	Class int    `json:"class"`
	Age   int    `json:"age"`
	Phone string `json:"phone" gorm:"type:varchar(40)"`
}

type Teacher struct {
	Id          int    `json:"id" gorm:"primary_key; auto_increment"`
	FirstName   string `json:"firstname" gorm:"type:varchar(40); not null"`
	LastName    string `json:"lastname" gorm:"type:varchar(40)"`
	PhoneNumber string `json:"phone number" gorm:"type:varchar(40)"`
}

type StudentAttendance struct {
	StudentId int `json:"roll" gorm:"primary_key"`
	PunchIn  time.Time `json:"punchin"`
	PunchOut time.Time `json:"punchout"`
}