package controller

import (
	"ichimonApi/model"
	"ichimonApi/usecase"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
	CsrfToken(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

// SignUp godoc
// @Summary ユーザー登録
// @Description 新しいユーザーを登録します。
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.UserCreateRequest true "Register User"
// @Success 200 {object} model.UserResponse
// @Router /signup [post]
func (uc *userController) SignUp(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, userRes)
}

// LogIn godoc
// @Summary ログイン
// @Description クッキーにログインセッションを記録する。
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.UserCreateRequest true "Login Credentials"
// @Success 200 "OK - Successfully authenticated"
// @Failure 400 {object} string "Bad Request - Invalid request format or parameters"
// @Failure 401 {object} string "Unauthorized - Invalid login credentials"
// @Router /login [post]
func (uc *userController) LogIn(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tokenString, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	// cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

// LogOut godoc
// @Summary ログアウト
// @Description クッキーに格納されているログインセッションを削除する。
// @Tags users
// @Accept json
// @Produce json
// @Success 204 "No Content - Successfully logged out"
// @Router /logout [get]
func (uc *userController) LogOut(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

// CsrfToken godoc
// @Summary CSRFトークンを取得する。
// @Description CSRFトークンを取得する。
// @Tags security
// @Accept json
// @Produce json
// @Success 200 {object} model.CsrfTokenResponse "OK - CSRF token provided"
// @Router /csrf [get]
func (uc *userController) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}