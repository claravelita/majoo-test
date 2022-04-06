package merchant

import (
	"github.com/claravelita/majoo-test/repository"
	"github.com/claravelita/majoo-test/usecase"
)

type merchantImplementation struct {
	repo repository.MerchantRepository
}

func NewMerchantImplementation(repo repository.MerchantRepository) usecase.MerchantUsecase {
	return &merchantImplementation{
		repo: repo,
	}
}
