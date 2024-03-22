package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

//IUserControllerはUserControllerのインターフェース
type IUserController interface {
	SingUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
}

//UserControllerはユーザーのリクエストを処理する
type UserController struct {
	uu usecase.IUserUsecase
}

//NewUserControllerはUserControllerのポインタを返す
func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &UserController{uu}
}

//SingUpはユーザー登録を行う
func (uc *UserController) SingUp(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userRes, err := uc.uu.SignUp(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, userRes)
}

func (uc *UserController) LogIn(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tokenString, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
}

func (uc *UserController) LogOut(c echo.Context) error {
}