package repository

import (
	"github.com/claravelita/majoo-test/common/dto"
	"github.com/claravelita/majoo-test/domain"
)

//go:generate mockgen -destination=../mocks/repository/mock_repository.go -package=mock_repository -source=repository_interface.go

type (
	UserReposiory interface {
		FetchUserByUsername(username string) (*domain.User, error)
		CreateUser(user domain.User) (domain.User, error)
	}

	MerchantRepository interface {
		CreateMerchant(merchant domain.Merchant) (domain.Merchant, error)
		FetchMerchantByUser(userID int64) (*domain.Merchant, error)
		FetchMerchantOmzet(id, start, end string) (resp []dto.MerchantOmzetRepoResponse, err error)
		FetchMerchantOmzetPagination(id, start, end, offset, limit string) (resp []dto.MerchantOmzetRepoResponse, totalData int64, err error)
	}

	OutletRepository interface {
		FetchOutletOmzet(id, start, end string) (resp []dto.OutletOmzetRepoResponse, err error)
		FetchOutletByIDAndUser(userID, id int64) (*domain.Outlet, error)
		FetchOutletOmzetByID(merchantID, id, start, end, offset, limit string) (resp []dto.OutletOmzetRepoResponse, totalData int64, err error)
	}
)
