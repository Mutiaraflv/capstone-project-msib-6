package controller

import (
	"net/http"

	"github.com/haerul-umam/capstone-project-mikti/model/web"
	"github.com/haerul-umam/capstone-project-mikti/service"
	"github.com/labstack/echo/v4"
)

type AuthControllerImpl struct {
	userService service.UserService
}

func NewAuthController(service service.UserService) *AuthControllerImpl {
	return &AuthControllerImpl{service}
}

func (controller *AuthControllerImpl) Login(e echo.Context) error {
	user := new(web.UserLoginRequest)

	if err := e.Bind(&user); err != nil {
		return e.JSON(400, web.ResponseToClient(400, err.Error(), nil))
	}

	if err := e.Validate(user); err != nil {
		return err
	}

	login, err := controller.userService.LoginUser(user.Email, user.Password)

	if err != nil {
		return e.JSON(
			400, web.ResponseToClient(400, err.Error(), nil),
		)
	}

	return e.JSON(200, web.ResponseToClient(200, "Sukses login", login))
}

func (controller *AuthControllerImpl) Register(e echo.Context) error {
	user := new(web.UserRegisterRequest)

	if err := e.Bind(user); err != nil {
		return e.JSON(http.StatusBadRequest, web.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := e.Validate(user); err != nil {
		return err
	}

	saveUser, errSaveUser := controller.userService.SaveUser(*user)

	if errSaveUser != nil {
		return e.JSON(http.StatusBadRequest, web.ResponseToClient(http.StatusBadRequest, errSaveUser.Error(), nil))
	}

	return e.JSON(http.StatusOK, web.ResponseToClient(http.StatusOK, "sukses register", saveUser))
}
