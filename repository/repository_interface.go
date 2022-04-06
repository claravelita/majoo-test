package repository

import "github.com/claravelita/majoo-test/domain"

//go:generate mockgen -destination=../mocks/repository/mock_repository.go -package=mock_repository -source=repository_interface.go

type (
	UserReposiory interface {
		FetchUserByUsername(username string) (*domain.User, error)
		CreateUser(user domain.User) (domain.User, error)
	}

	MerchantRepository interface {
		CreateMerchant(merchant domain.Merchant) (domain.Merchant, error)
	}
)
