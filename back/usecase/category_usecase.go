package usecase

import (
	"ichimonApi/model"
	"ichimonApi/repository"
	"ichimonApi/utility"
	"ichimonApi/validator"
)

type ICategoryUsecase interface {
	FindById(id uint) (model.CategoryResponse, error)
	FindAll() ([]model.CategoryResponse, error)
	CreateCategory(category model.Category) (model.CategoryResponse, error)
}

type categoryUsecase struct {
	cr repository.ICategoryRepository
	cv validator.ICategoryValidator
}

func NewCategoryUsecase(cr repository.ICategoryRepository,cv validator.ICategoryValidator) ICategoryUsecase {
	return &categoryUsecase{cr,cv}
}

func (cu *categoryUsecase) FindById(categoryId uint) (model.CategoryResponse, error) {
	category := model.Category{}
	if err := cu.cr.FindByCategoryId(&category, categoryId); err != nil {
		return model.CategoryResponse{}, err
	}
	resCategory := model.CategoryResponse{}
	utility.CopyFields(category, &resCategory)
	return resCategory, nil
}

func (cu *categoryUsecase) FindAll() ([]model.CategoryResponse, error) {
	categories := []model.Category{}
	if err := cu.cr.FindAllCategories(&categories); err != nil {
		return nil, err
	}
	resCategories := []model.CategoryResponse{}
	for _, category := range categories {
		resCategory := model.CategoryResponse{}
		//resQuestionにquestionの値が反映されていることを期待
		utility.CopyFields(category, &resCategory)
		resCategories = append(resCategories, resCategory)
	}
	return resCategories, nil
}

func (cu *categoryUsecase) CreateCategory(category model.Category) (model.CategoryResponse, error) {
	if err := cu.cv.CategoryValidate(category); err != nil {
		return model.CategoryResponse{}, err
	}
	if err := cu.cr.CreateCategory(&category); err != nil {
		return model.CategoryResponse{}, err
	}
	resCategory := model.CategoryResponse{}
	utility.CopyFields(category, &resCategory)
	return resCategory, nil
}
