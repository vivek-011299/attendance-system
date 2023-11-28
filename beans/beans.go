package beans

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Student struct {
	*gorm.DB
	Id    int    `json:"roll" gorm:"primary_key; auto_increment"`
	Name  string `json:"name" gorm:"type:varchar(40); not null"`
	Class int    `json:"class"`
	Age   int    `json:"age"`
	Phone string `json:"phone" gorm:"type:varchar(40)"`
}

type Teacher struct {
	*gorm.DB
	Id          int    `json:"id" gorm:"primary_key; auto_increment"`
	FirstName   string `json:"firstname" gorm:"type:varchar(40); not null"`
	LastName    string `json:"lastname" gorm:"type:varchar(40)"`
	PhoneNumber string `json:"phone number" gorm:"type:varchar(40)"`
}

type StudentAttendance struct {
	*gorm.DB
	StudentId int `json:"roll" gorm:"type:integer; primary_key"`
	PunchIn  time.Time `json:"punchin" gorm:"type:TIME"`
	PunchOut time.Time `json:"punchout" gorm:"type:TIME"`
}

var Db *gorm.DB
