package model

import (
	"database/sql"
	"gin-demo/pkg/utils"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int64 `gorm:"primarykey"`
	Name      string
	Age       int
	Sex       int
	Courses   []Course `gorm:"many2many:student_course;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}

func (s *User) BeforeCreate(db *gorm.DB) error {
	id, err := utils.GenerateSnowflakeId()
	if err != nil {
		return err
	}
	s.ID = id
	return nil
}
