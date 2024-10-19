package model

type Student struct {
	BaseModel
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Sex     int      `json:"sex"`
	Courses []Course `gorm:"many2many:student_course;"`
}

//func (s *Student) BeforeCreate(db *gorm.DB) error {
//	id, err := utils.GenerateSnowflakeId()
//	if err != nil {
//		return err
//	}
//	s.ID = id
//	return nil
//}
