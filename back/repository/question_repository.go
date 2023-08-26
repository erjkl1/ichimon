package repository

import (
	"ichimonApi/model"

	"gorm.io/gorm"
)


type IQuestionRepository interface {
	FindById(question *model.Question, id uint) error
	FindBySubCategoryId(question *model.Question, subCategoryId uint) error
	FindByCategoryId(question *model.Question, subCategoryId uint) error
	Save(question *model.Question) error
}

type questionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) IQuestionRepository {
	return &questionRepository{db}
}

func (qr *questionRepository) FindById(question *model.Question, id uint) error {
	if err := qr.db.Where("id=?", id).First(question).Error; err != nil {
		return err
	}
	return nil
}

func (qr *questionRepository) FindBySubCategoryId(question *model.Question, subCategoryId uint) error {
	if err := qr.db.Where("sub_category_id=?", subCategoryId).First(question).Error; err != nil {
		return err
	}
	return nil
}

// subcategoryをJOINしてその関係性で検索する
func (qr *questionRepository) FindByCategoryId(question *model.Question, categoryId uint) error {
	if err := qr.db.Where("category_id=?", categoryId).First(question).Error; err != nil {
		return err
	}
	return nil
}

func (qr *questionRepository) Save(question *model.Question) error {
	if err := qr.db.Create(question).Error; err != nil {
		return err
	}
	return nil
}