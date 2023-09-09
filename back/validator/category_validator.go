package validator

import (
	"ichimonApi/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ICategoryValidator interface {
	CategoryValidate(category model.Category) error
}

type categoryValidator struct{}

func NewCategoryValidator() ICategoryValidator {
	return &categoryValidator{}
}

func (tv *categoryValidator) CategoryValidate(category model.Category) error {
	return validation.ValidateStruct(&category,
		validation.Field(
			&category.SubCategoryId,
			validation.Required.Error("sub_category_id is required"),
		),
		validation.Field(
			&category.Title,
			validation.Required.Error("title is required"),
			validation.RuneLength(1, 50).Error("limited max 50 char"),
		),
		validation.Field(
			&category.Statement,
			validation.Required.Error("statement is required"),
			validation.RuneLength(1, 300).Error("limited max 300 char"),
		),
		validation.Field(
			&category.Selection1,
			validation.Required.Error("selection_1 is required"),
			validation.RuneLength(1, 50).Error("limited max 50 char"),
		),
		validation.Field(
			&category.Selection2,
			validation.Required.Error("selection_2 is required"),
			validation.RuneLength(1, 50).Error("limited max 50 char"),
		),
		validation.Field(
			&category.Selection3,
			validation.Required.Error("selection_3 is required"),
			validation.RuneLength(1, 50).Error("limited max 50 char"),
		),
		validation.Field(
			&category.Selection4,
			validation.Required.Error("selection_4 is required"),
			validation.RuneLength(1, 50).Error("limited max 50 char"),
		),
		validation.Field(
			&category.Answer,
			validation.Required.Error("answer is required"),
		),
		validation.Field(
			&category.Description,
			validation.Required.Error("description is required"),
			validation.RuneLength(1, 300).Error("limited max 300 char"),
		),
	)
}