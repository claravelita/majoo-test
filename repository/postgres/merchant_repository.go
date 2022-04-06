package postgres

import (
	"github.com/claravelita/majoo-test/domain"
	"github.com/claravelita/majoo-test/repository"
	"gorm.io/gorm"
)

type merchantRepository struct {
	db *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) repository.MerchantRepository {
	return merchantRepository{
		db,
	}
}

func (r merchantRepository) CreateMerchant(merchant domain.Merchant) (domain.Merchant, error) {
	err := r.db.Create(&merchant).Error
	if err != nil {
		return domain.Merchant{}, err
	}

	return merchant, nil
}
