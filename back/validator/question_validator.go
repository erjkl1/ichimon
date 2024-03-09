package validator

import (
	"ichimonApi/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IQuestionValidator interface {
	QuestionValidate(question model.Question) error
}

type questionValidator struct{}

func NewQuestionValidator() IQuestionValidator {
	return &questionValidator{}
}

func (tv *questionValidator) QuestionValidate(question model.Question) error {
	return validation.ValidateStruct(&question,
		validation.Field(
			&question.SubCategoryId,
			validation.Required.Error("sub_category_id is required"),
		),
		validation.Field(
			&question.Title,
			validation.Required.Error("title is required"),
			validation.RuneLength(1, 50).Error("limited max 50 char"),
		),
		validation.Field(
			&question.Statement,
			validation.Required.Error("statement is required"),
			validation.RuneLength(1, 300).Error("limited max 300 char"),
		),
		validation.Field(
			&question.Selection1,
			validation.Required.Error("selection_1 is required"),
			validation.RuneLength(1, 50).Error("limited max 50 char"),
		),
		validation.Field(
			&question.Selection2,
			validation.Required.Error("selection_2 is required"),
			validation.RuneLength(1, 50).Error("limited max 50 char"),
		),
		validation.Field(
			&question.Selection3,
			validation.Required.Error("selection_3 is required"),
			validation.RuneLength(1, 50).Error("limited max 50 char"),
		),
		validation.Field(
			&question.Selection4,
			validation.Required.Error("selection_4 is required"),
			validation.RuneLength(1, 50).Error("limited max 50 char"),
		),
		validation.Field(
			&question.Answer,
			validation.Required.Error("answer is required"),
		),
		validation.Field(
			&question.Description,
			validation.Required.Error("description is required"),
			validation.RuneLength(1, 300).Error("limited max 300 char"),
		),
	)
}