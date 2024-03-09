package repository

import (
	"ichimonApi/model"

	"gorm.io/gorm"
)


type IGradeRepository interface {
	FindById(grade *model.Grade, id uint) error
	Save(grade *model.Grade) error 
}

type gradeRepository struct {
	db *gorm.DB
}

func NewGradeRepository(db *gorm.DB) IGradeRepository {
	return &gradeRepository{db}
}

func (gr *gradeRepository) FindById(grade *model.Grade, id uint) error {
	if err := gr.db.Where("id=?", id).First(grade).Error; err != nil {
		return err
	}
	return nil
}

func (gr *gradeRepository) Save(grade *model.Grade) error {
	if err := gr.db.Create(grade).Error; err != nil {
		return err
	}
	return nil
}