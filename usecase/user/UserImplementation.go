package user

import (
	"github.com/claravelita/majoo-test/common/helper"
	"github.com/claravelita/majoo-test/repository"
	"github.com/claravelita/majoo-test/usecase"
)

type userImplementation struct {
	repo repository.UserReposiory
	auth helper.AuthHelper
}

func NewUserImplementation(repo repository.UserReposiory, auth helper.AuthHelper) usecase.UserUsecase {
	return &userImplementation{
		repo: repo,
		auth: auth,
	}
}
