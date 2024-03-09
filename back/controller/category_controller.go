package controller

import (
	"ichimonApi/model"
	"ichimonApi/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ICategoryController interface {
	FindById(c echo.Context) error
	FindAll(c echo.Context) error
	CreateCategory(c echo.Context) error
	CreateSubCategory(c echo.Context) error
}

type categoryController struct {
	cu usecase.ICategoryUsecase
}

func NewCategoryController(cu usecase.ICategoryUsecase) ICategoryController {
	return &categoryController{cu}
}

// FindById godoc
// @Summary カテゴリIDによる検索
// @Description 指定されたIDを持つカテゴリを検索し、結果を返します。
// @Tags category
// @Accept  json
// @Produce  json
// @Param categoryId path int true "Category ID"
// @Success 200 {object} model.CategoryResponse
// @Router /categories/{categoryId} [get]
func (cc *categoryController) FindById(c echo.Context) error {

	categoryId, _ := strconv.Atoi(c.Param("categoryId"))
	
	resCategory, err := cc.cu.FindById(uint(categoryId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resCategory)
}

//TODO sub_categoryの関連がおかしい？取れない。
// FindAll godoc
// @Summary すべてのカテゴリの検索
// @Description 登録されているすべてのカテゴリを検索し、結果を返します。
// @Tags category
// @Accept  json
// @Produce  json
// @Success 200 {array} model.CategoryResponse
// @Router /categories [get]
func (cc *categoryController) FindAll(c echo.Context) error {	
	categories, err := cc.cu.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, categories)
}

//TODO view_order指定できない問題
// CreateCategory godoc
// @Summary カテゴリの作成
// @Description 新しいカテゴリを作成します。
// @Tags category
// @Accept  json
// @Produce  json
// @Param category body model.CreateCategoryRequest true "Create Category"
// @Success 201 {object} model.CategoryResponse
// @Router /categories [post]
func (cc *categoryController) CreateCategory(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	createdUserId := claims["user_id"]

	category := model.Category{}
	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// category.UpdatedUserId = uint(createdUserId.(float64))
	category.CreatedUserId = uint(createdUserId.(float64))
	categoryResponse, err := cc.cu.CreateCategory(category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, categoryResponse)
}

// CreateSubCategory godoc
// @Summary サブカテゴリの作成
// @Description 新しいサブカテゴリを作成します。
// @Tags category
// @Accept  json
// @Produce  json
// @Param subCategory body model.CreateSubCategoryRequest true "Create SubCategory"
// @Success 201 {object} model.SubCategoryResponse
// @Router /subcategories [post]
func (cc *categoryController) CreateSubCategory(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	createdUserId := claims["user_id"]

	subCategory := model.SubCategory{}
	if err := c.Bind(&subCategory); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// subCategory.UpdatedUserId = uint(createdUserId.(float64))
	subCategory.CreatedUserId = uint(createdUserId.(float64))
	categoryResponse, err := cc.cu.CreateSubCategory(subCategory)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, categoryResponse)
}