package model

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Students    []Student `gorm:"many2many:student_course;"`
}
