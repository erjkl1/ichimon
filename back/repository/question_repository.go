package repository

import (
	"ichimonApi/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)


type IQuestionRepository interface {
	FindById(question *model.Question, id uint) error
	FindAllBySubCategoryId(questions *[]model.Question, subCategoryId uint) error
	FindAllBySubCategoryIdAndGradeUserId(questions *[]model.Question, subCategoryId uint, gradeUserId uint) error
	FindByCategoryId(question *model.Question, subCategoryId uint) error
	FindAllByCreatedUserId(questions *[]model.Question, createdUserId uint) error
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

// subcategoryをJOINしてその関係性で検索する
func (qr *questionRepository) FindByCategoryId(questions *model.Question, categoryId uint) error {
	if err := qr.db.Where("category_id=?", categoryId).First(questions).Error; err != nil {
		return err
	}
	return nil
}

func (qr *questionRepository) FindAllByCreatedUserId(questions *[]model.Question, createdUserId uint) error {
	if err := qr.db.Where("created_user_id=?", createdUserId).Find(questions).Error; err != nil {
		return err
	}
	return nil
}

//サブカテゴリIDで絞り込み問題に紐づく全てのテーブルの関連を取得
func (qr *questionRepository) FindAllBySubCategoryId(questions *[]model.Question, subCategoryId uint) error {
	if err := qr.db.Preload("SubCategory.Category").Preload(clause.Associations).Where("sub_category_id=?", subCategoryId).Find(questions).Error; err != nil {
		return err
	}
	return nil
}

//サブカテゴリIDと成績をユーザーIDで絞り込み問題に紐づく全てのテーブルの関連を取得
func (qr *questionRepository) FindAllBySubCategoryIdAndGradeUserId(questions *[]model.Question, subCategoryId uint, gradeUserId uint) error {
	if err := qr.db.Preload("SubCategory.Category").Preload(clause.Associations).Preload("Grade", "user_id = ?" , gradeUserId).Where("sub_category_id=?", subCategoryId).Order("created_at DESC").Find(questions).Error; err != nil {
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