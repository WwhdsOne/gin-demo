package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name    string
	Age     int
	Sex     int
	Courses []Course `gorm:"many2many:student_course;"`
}
