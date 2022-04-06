package controller

import (
	"github.com/claravelita/majoo-test/common/dto"
	"github.com/claravelita/majoo-test/common/models"
	"github.com/claravelita/majoo-test/usecase"
	"github.com/labstack/echo/v4"
)

type userController struct {
	usecase usecase.UserUsecase
}

func NewUserController(usecase usecase.UserUsecase) *userController {
	return &userController{usecase: usecase}
}

func (c *userController) Route(group *echo.Group) {
	group.POST("/register", c.Register)
	group.POST("/log-in", c.Login)
}

func (c *userController) Register(ctx echo.Context) error {
	request := dto.UserRequest{}

	err := models.ValidateRequest(ctx, &request)
	if err != nil {
		return err
	}

	responses, err := c.usecase.CreateUserCommand(request)

	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

func (c *userController) Login(ctx echo.Context) error {
	request := dto.UserLoginRequest{}

	err := models.ValidateRequest(ctx, &request)
	if err != nil {
		return err
	}

	responses, err := c.usecase.LoginUserCommand(request)

	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}
