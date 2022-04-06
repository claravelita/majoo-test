package outlet

import (
	"github.com/claravelita/majoo-test/repository"
	"github.com/claravelita/majoo-test/usecase"
)

type outletImplementation struct {
	repo         repository.OutletRepository
	repoMerchant repository.MerchantRepository
}

func NewOutletImplementation(repo repository.OutletRepository, repoMerchant repository.MerchantRepository) usecase.OutletUsecase {
	return &outletImplementation{
		repo:         repo,
		repoMerchant: repoMerchant,
	}
}
