package usecase

import (
	"ichimonApi/model"
	"ichimonApi/repository"
	"ichimonApi/validator"
)

type ICategoryUsecase interface {
	FindById(id uint) (model.Category, error)
	FindAll() ([]model.Category, error)
	CreateCategory(category model.Category) (model.Category, error)
}

type categoryUsecase struct {
	cr repository.ICategoryRepository
	cv validator.ICategoryValidator
}

func NewCategoryUsecase(cr repository.ICategoryRepository,cv validator.ICategoryValidator) ICategoryUsecase {
	return &categoryUsecase{cr,cv}
}

func (cu *categoryUsecase) FindById(categoryId uint) (model.Category, error) {
	category := model.Category{}
	if err := cu.cr.FindByCategoryId(&category, categoryId); err != nil {
		return model.Category{}, err
	}
	return category, nil
}

func (cu *categoryUsecase) FindAll() ([]model.Category, error) {
	categories := []model.Category{}
	if err := cu.cr.FindAllCategories(&categories); err != nil {
		return nil, err
	}
	return categories, nil
}

func (cu *categoryUsecase) CreateCategory(category model.Category) (model.Category, error) {
	if err := cu.cv.CategoryValidate(category); err != nil {
		return model.Category{}, err
	}
	if err := cu.cr.CreateCategory(&category); err != nil {
		return model.Category{}, err
	}
	return category, nil
}
