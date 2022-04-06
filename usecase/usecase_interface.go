package usecase

import (
	"github.com/claravelita/majoo-test/common/dto"
	"github.com/claravelita/majoo-test/common/models"
)

type (
	UserUsecase interface {
		CreateUserCommand(request dto.UserRequest) (models.JSONResponses, error)
		LoginUserCommand(request dto.UserLoginRequest) (models.JSONResponses, error)
	}
)
