package validator

import (
	"ichimonApi/model"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ISubCategoryValidator interface {
	SubCategoryValidate(subCategory model.SubCategory) error
}

type subCategoryValidator struct{}

func NewSubCategoryValidator() ISubCategoryValidator {
	return &subCategoryValidator{}
}

func (tv *subCategoryValidator) SubCategoryValidate(subCategory model.SubCategory) error {
	return validation.ValidateStruct(&subCategory,
		validation.Field(
			&subCategory.CreatedUserId,
			validation.Required.Error("create_user_id is required"),
		),
		validation.Field(
			&subCategory.CategoryId,
			validation.Required.Error("subCategory_id is required"),
		),
		validation.Field(
			&subCategory.SubCategoryName,
			validation.Required.Error("subCategory_name is required"),
			validation.RuneLength(1, 50).Error("limited max 50 char"),
		),
		validation.Field(
			&subCategory.ViewOrder,
			is.Int.Error("view_order is required"),
		),
	)
}