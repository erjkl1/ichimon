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
			&question.Title,
			validation.Required.Error("title is required"),
			validation.RuneLength(1, 10).Error("limited max 10 char"),
		),
	)
}