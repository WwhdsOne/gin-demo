package service

import (
	"gin-demo/global"
	"gin-demo/model"
)

func AddStudent(s *model.Student) error {
	tx := global.DB.Create(&s)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func DeleteStudent(id int) error {
	tx := global.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// 先查找学生信息
	var student model.Student
	if err := tx.Where("id = ?", id).First(&student).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除学生与课程的关系
	if err := tx.Model(&student).Association("Courses").Clear(); err != nil {
		tx.Rollback()
		return err
	}

	// 删除学生信息
	if err := tx.Delete(&student).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func UpdateStudent(s *model.Student) error {
	// 开始事务
	tx := global.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// 更新学生信息
	if err := tx.Model(&model.Student{}).Where("id = ?", s.BaseModel.ID).
		Updates(&s).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 替换关联关系
	if err := tx.Model(s).Association("Courses").Replace(&s.Courses); err != nil {
		tx.Rollback()
		return err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func GetStudent(id int) (*model.Student, error) {
	var s model.Student
	//global.LOG.Info("GetStudent")
	tx := global.DB.Model(&model.Student{}).
		Where("id = ?", id).
		Preload("Courses").
		First(&s)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &s, nil
}
