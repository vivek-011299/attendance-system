package beans

import (
	"github.com/jinzhu/gorm"
)

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
	PhoneNumber string `json:"phonenumber" gorm:"type:varchar(40)"`
}

type StudentAttendance struct {
	StudentId int `json:"roll"`
	PunchIn  string `json:"punchin" gorm:"type:varchar(40)"`
	PunchOut string `json:"punchout" gorm:"type:varchar(40)"`
}

type MessageBox struct {
	Message string `json:"message"`
}

var Db *gorm.DB