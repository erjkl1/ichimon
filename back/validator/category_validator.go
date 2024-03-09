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
			&category.CreatedUserId,
			validation.Required.Error("create_user_id is required"),
		),
		validation.Field(
			&category.CategoryName,
			validation.Required.Error("category_name is required"),
			validation.RuneLength(1, 50).Error("limited max 50 char"),
		),
		validation.Field(
			&category.ViewOrder,
			validation.Required.Error("view_order is required"),
		),
	)
}