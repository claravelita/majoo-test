package user

import (
	"github.com/claravelita/majoo-test/common/command"
	"github.com/claravelita/majoo-test/common/dto"
	"github.com/claravelita/majoo-test/common/models"
	"github.com/claravelita/majoo-test/domain"
)

func (u userImplementation) CreateUserCommand(request dto.UserRequest) (models.JSONResponses, error) {
	password, err := u.auth.HashPassword(request.Password)
	if err != nil {
		return command.InternalServerResponses("Internal Server Error"), err
	}

	checkUsername, err := u.repo.FetchUserByUsername(request.UserName)
	if checkUsername != nil {
		return command.BadRequestResponses("Username is Already Registered"), nil
	}
	if err != nil {
		return command.InternalServerResponses("Internal Server Error"), err
	}

	requestUser := domain.User{
		Name:     request.Name,
		UserName: request.UserName,
		Password: password,
		Merchants: domain.Merchant{
			MerchantName: request.MerchantName,
		},
	}
	newUser, err := u.repo.CreateUser(requestUser)
	if err != nil {
		return command.InternalServerResponses("Internal Server Error"), err
	}

	token, tokenErr := u.auth.CreateJWTTokenLogin(string(newUser.ID), newUser.Name, newUser.UserName)
	if tokenErr != nil {
		return command.InternalServerResponses("Internal Server Error"), err
	}

	data := dto.UserAccessResponse{
		Name:         newUser.Name,
		UserName:     newUser.UserName,
		MerchantName: request.MerchantName,
		AccessToken:  "Bearer " + token,
	}

	return command.SuccessResponses(data), nil
}
