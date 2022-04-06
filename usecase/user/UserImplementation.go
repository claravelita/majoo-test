package user

import (
	"github.com/claravelita/majoo-test/repository"
	"github.com/claravelita/majoo-test/usecase"
)

type userImplementation struct {
	repo repository.UserReposiory
}

func NewUserImplementation(repo repository.UserReposiory) usecase.UserUsecase {
	return &userImplementation{
		repo: repo,
	}
}
