package repository

import (
	"ichimonApi/model"

	"gorm.io/gorm"
)


type ICategoryRepository interface {
	FindByCategoryId(category *model.Category, subCategoryId uint) error
	FindBySubCategoryId(category *model.SubCategory, subCategoryId uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &categoryRepository{db}
}

func (cr *categoryRepository) FindByCategoryId(category *model.Category, categoryId uint) error {
	if err := cr.db.Where("category_id=?", categoryId).First(category).Error; err != nil {
		return err
	}
	return nil
}

func (cr *categoryRepository) FindBySubCategoryId(category *model.SubCategory, subCategoryId uint) error {
	if err := cr.db.Where("sub_category_id=?", subCategoryId).First(category).Error; err != nil {
		return err
	}
	return nil
}