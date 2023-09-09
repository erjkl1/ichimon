package usecase

import (
	"ichimonApi/model"
	"ichimonApi/repository"
	"ichimonApi/utility"
	"ichimonApi/validator"
)
type IQuestionUsecase interface {

	FindById(id uint) (model.QuestionResponse, error)
	FindAll() ([]model.QuestionResponse, error)
	FindAllByCategoryId(categoryId uint) ([]model.QuestionResponse, error)
	FindAllBySubCategoryId(subCategoryId uint) ([]model.QuestionResponse, error)
	FindAllByCreatedUserId(createdUserId uint) ([]model.QuestionResponse, error)
	CreateQuestion(model.Question) (model.QuestionResponse, error)
	// CreateQuestion(question model.Question) (model.QuestionResponse, error)
	// UpdateQuestion(question model.Question, userId uint, questionId uint) (model.QuestionResponse, error)
}

type questionUsecase struct {
	tr repository.IQuestionRepository
	tv validator.IQuestionValidator
}

func NewQuestionUsecase(qr repository.IQuestionRepository,tv validator.IQuestionValidator) IQuestionUsecase {
	return &questionUsecase{qr, tv}
}

func (qu *questionUsecase) FindById(questionId uint) (model.QuestionResponse, error) {
	question := model.Question{}
	if err := qu.tr.FindById(&question, questionId); err != nil {
		return model.QuestionResponse{}, err
	}
	resQuestion := model.QuestionResponse{}
	utility.CopyFields(question, &resQuestion)
	return resQuestion, nil
}

func (qu *questionUsecase) FindAll() ([]model.QuestionResponse, error) {
	questions := []model.Question{}
	if err := qu.tr.FindAll(&questions); err != nil {
		return nil, err
	}
	resQuestions := []model.QuestionResponse{}
	for _, question := range questions {
		resQuestion := model.QuestionResponse{}
		//resQuestionにquestionの値が反映されていることを期待
		utility.CopyFields(question, &resQuestion)
		resQuestions = append(resQuestions, resQuestion)
	}
	return resQuestions, nil
}


func (qu *questionUsecase) FindAllByCategoryId(categoryId uint) ([]model.QuestionResponse, error) {
	questions := []model.Question{}
	if err := qu.tr.FindAllByCategoryId(&questions, categoryId); err != nil {
		return nil, err
	}
	resQuestions := []model.QuestionResponse{}
	for _, question := range questions {
		resQuestion := model.QuestionResponse{}
		//resQuestionにquestionの値が反映されていることを期待
		utility.CopyFields(question, &resQuestion)
		resQuestions = append(resQuestions, resQuestion)
	}
	return resQuestions, nil
}

func (qu *questionUsecase) FindAllBySubCategoryId(subCategoryId uint) ([]model.QuestionResponse, error) {
	questions := []model.Question{}
	if err := qu.tr.FindAllBySubCategoryId(&questions, subCategoryId); err != nil {
		return nil, err
	}
	resQuestions := []model.QuestionResponse{}
	for _, question := range questions {
		resQuestion := model.QuestionResponse{}
		//resQuestionにquestionの値が反映されていることを期待
		utility.CopyFields(question, &resQuestion)
		resQuestions = append(resQuestions, resQuestion)
	}
	return resQuestions, nil
}

func (qu *questionUsecase) FindAllByCreatedUserId(createdUserId uint) ([]model.QuestionResponse, error) {
	questions := []model.Question{}
	if err := qu.tr.FindAllByCreatedUserId(&questions, createdUserId); err != nil {
		return nil, err
	}
	resQuestions := []model.QuestionResponse{}

	for _, question := range questions {
		resQuestion := model.QuestionResponse{}
		//resQuestionにquestionの値が反映されていることを期待
		utility.CopyFields(question, &resQuestion)
		resQuestions = append(resQuestions, resQuestion)
	}
	return resQuestions, nil
}

func (qu *questionUsecase) CreateQuestion(question model.Question) (model.QuestionResponse, error) {
	if err := qu.tv.QuestionValidate(question); err != nil {
		return model.QuestionResponse{}, err
	}
	if err := qu.tr.CreateQuestion(&question); err != nil {
		return model.QuestionResponse{}, err
	}
	resQuestion := model.QuestionResponse{}
	utility.CopyFields(question, &resQuestion)
	return resQuestion, nil
}

// func (qu *questionUsecase) UpdateQuestion(question model.Question, userId uint, questionId uint) (model.QuestionResponse, error) {
// 	if err := qu.tv.QuestionValidate(question); err != nil {
// 		return model.QuestionResponse{}, err
// 	}
// 	if err := qu.tr.UpdateQuestion(&question, userId, questionId); err != nil {
// 		return model.QuestionResponse{}, err
// 	}
// 	resQuestion := model.QuestionResponse{
// 		ID:        question.ID,
// 		Title:     question.Title,
// 		CreatedAt: question.CreatedAt,
// 		UpdatedAt: question.UpdatedAt,
// 	}
// 	return resQuestion, nil
// }

// func (qu *questionUsecase) DeleteQuestion(userId uint, questionId uint) error {
// 	if err := qu.tr.DeleteQuestion(userId, questionId); err != nil {
// 		return err
// 	}
// 	return nil
// }