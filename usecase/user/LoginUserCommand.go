package user

import (
	"github.com/claravelita/majoo-test/common/command"
	"github.com/claravelita/majoo-test/common/dto"
	"github.com/claravelita/majoo-test/common/models"
	"strconv"
)

func (u userImplementation) LoginUserCommand(request dto.UserLoginRequest) (models.JSONResponses, error) {
	findUser, err := u.repo.FetchUserByUsername(request.UserName)
	if findUser == nil {
		return command.BadRequestResponses("Couldn't found your account"), nil
	}
	if err != nil {
		return command.InternalServerResponses("Internal Server Error"), err
	}

	checkPassword := u.auth.CheckPasswordHash(request.Password, findUser.Password)
	if checkPassword == true {
		token, tokenErr := u.auth.CreateJWTTokenLogin(strconv.FormatInt(findUser.ID, 10), findUser.Name, findUser.UserName)
		if tokenErr != nil {
			return command.InternalServerResponses("Internal Server Error"), err
		}
		data := dto.UserAccessResponse{
			Name:         findUser.Name,
			UserName:     findUser.UserName,
			MerchantName: findUser.Merchants.MerchantName,
			AccessToken:  "Bearer " + token,
		}
		return command.SuccessResponses(data), nil
	}

	return command.BadRequestResponses("Email or Password is incorrect"), nil
}
