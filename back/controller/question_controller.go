package controller

import (
	"ichimonApi/model"
	"ichimonApi/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type IQuestionController interface {
	FindById(c echo.Context) error
	FindAll(c echo.Context) error
	FindAllByCategoryId(c echo.Context) error
	FindAllBySubCategoryId(c echo.Context) error
	FindAllByCreatedUserId(c echo.Context) error
	CreateQuestion(c echo.Context) error
	// CsrfToken(c echo.Context) error
}

type questionController struct {
	qu usecase.IQuestionUsecase
}

func NewQuestionController(qu usecase.IQuestionUsecase) IQuestionController {
	return &questionController{qu}
}

// FindById godoc
// @Summary 問題IDによる検索
// @Description 指定されたIDを持つ問題を検索し、結果を返します。
// @Tags question
// @Accept  json
// @Produce  json
// @Param questionId path int true "Question ID"
// @Success 200 {object} model.QuestionResponse
// @Failure 500 {object} string "Internal Server Error"
// @Router /questions/{questionId} [get]
func (qc *questionController) FindById(c echo.Context) error {

	questionId, _ := strconv.Atoi(c.Param("questionId"))
	
	questionResponse, err := qc.qu.FindById(uint(questionId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, questionResponse)
}

// FindAll godoc
// @Summary すべての問題の検索
// @Description 登録されているすべての問題を検索し、結果を返します。
// @Tags question
// @Accept  json
// @Produce  json
// @Success 200 {array} model.QuestionResponse
// @Failure 500 {object} string "Internal Server Error"
// @Router /questions [get]
func (qc *questionController) FindAll(c echo.Context) error {
	questionResponses, err := qc.qu.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, questionResponses)
}


// FindAllByCategoryId godoc
// @Summary カテゴリIDに基づく問題の検索
// @Description 指定されたカテゴリIDに属する問題をすべて検索し、結果を返します。
// @Tags question
// @Accept  json
// @Produce  json
// @Param categoryId path int true "Category ID"
// @Success 200 {array} model.QuestionResponse
// @Failure 500 {object} string "Internal Server Error"
// @Router /questions/category/{categoryId} [get]
func (qc *questionController) FindAllByCategoryId(c echo.Context) error {

	categoryId, _ := strconv.Atoi(c.Param("categoryId"))
	
	questionResponses, err := qc.qu.FindAllByCategoryId(uint(categoryId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, questionResponses)
}


// FindAllBySubCategoryId godoc
// @Summary サブカテゴリIDに基づく問題の検索
// @Description 指定されたサブカテゴリIDに属する問題をすべて検索し、結果を返します。
// @Tags question
// @Accept  json
// @Produce  json
// @Param subCategoryId path int true "SubCategory ID"
// @Success 200 {array} model.QuestionResponse
// @Failure 500 {object} string "Internal Server Error"
// @Router /questions/subcategory/{subCategoryId} [get]
func (qc *questionController) FindAllBySubCategoryId(c echo.Context) error {

	categoryId, _ := strconv.Atoi(c.Param("subCategoryId"))
	
	questionResponses, err := qc.qu.FindAllBySubCategoryId(uint(categoryId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, questionResponses)
}

// FindAllByCreatedUserId godoc
// @Summary 作成者IDに基づく問題の検索
// @Description 指定された作成者IDで作成された問題をすべて検索し、結果を返します。
// @Tags question
// @Accept  json
// @Produce  json
// @Param createdUserId path int true "Created User ID"
// @Success 200 {array} model.QuestionResponse
// @Failure 500 {object} string "Internal Server Error"
// @Router /questions/user/{createdUserId} [get]
func (qc *questionController) FindAllByCreatedUserId(c echo.Context) error {

	categoryId, _ := strconv.Atoi(c.Param("createdUserId"))
	
	questionResponses, err := qc.qu.FindAllByCreatedUserId(uint(categoryId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, questionResponses)
}

// CreateQuestion godoc
// @Summary 新しい問題の作成
// @Description 新しい問題を作成し、データベースに保存します。
// @Tags question
// @Accept  json
// @Produce  json
// @Param question body model.Question true "Create Question"
// @Success 201 {object} model.QuestionResponse
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /questions [post]
func (qc *questionController) CreateQuestion(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	createdUserId := claims["user_id"]

	question := model.Question{}
	if err := c.Bind(&question); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	question.UpdatedUserId = uint(createdUserId.(float64))
	question.CreatedUserId = uint(createdUserId.(float64))
	questionResponse, err := qc.qu.CreateQuestion(question)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, questionResponse)
}

// func (qc *questionController) UpdateQuestion(c echo.Context) error {
// 	user := c.Get("user").(*jwt.Token)
// 	claims := user.Claims.(jwt.MapClaims)
// 	updatedUserId := claims["user_id"]

// 	// IDをjsonから取得？
// 	questionId, _ := strconv.Atoi(c.Param("questionId"))
// 	// 指定されたIDで見つからなければエラー
// 	questionResponse, err := qc.qu.FindById(uint(questionId))
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}


// 	question := model.Question{}
// 	if err := c.Bind(&question); err != nil {
// 		return c.JSON(http.StatusBadRequest, err.Error())
// 	}
// 	question.UpdatedUserId = uint(updatedUserId.(float64))
// 	questionResponse, err := qc.qu.UpdateQuestion(question)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}
// 	return c.JSON(http.StatusCreated, questionResponse)
// }

// func (qc *questionController) CsrfToken(c echo.Context) error {
// 	token := c.Get("csrf").(string)
// 	return c.JSON(http.StatusOK, echo.Map{
// 		"csrf_token": token,
// 	})
// }