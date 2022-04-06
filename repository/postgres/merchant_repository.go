package postgres

import (
	"errors"
	"github.com/claravelita/majoo-test/common/dto"
	"github.com/claravelita/majoo-test/domain"
	"github.com/claravelita/majoo-test/repository"
	"gorm.io/gorm"
	"strings"
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

func (r merchantRepository) FetchMerchantByUser(userID int64) (*domain.Merchant, error) {
	var merchantEntity domain.Merchant
	err := r.db.Preload("Outlets").
		Preload("Transactions").Where("user_id = ?", userID).First(&merchantEntity).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &merchantEntity, err
}

func (r merchantRepository) FetchMerchantOmzet(id, start, end string) (resp []dto.MerchantOmzetRepoResponse, err error) {
	var baseQuery, query string
	baseQuery = `SELECT day::date, (case when omzet is null then 0 else omzet end)
		FROM(
			SELECT day::date
				FROM generate_series
				('@start'::date
				,'@end'::date
				,'1 day'::interval) day
				) as d
			LEFT JOIN(SELECT T.created_at::date as day, sum(T.bill_total) as omzet
				FROM transactions AS T
				WHERE T.merchant_id = '@id'
				GROUP BY day) o USING (day);`

	query = strings.Replace(baseQuery, "@start", start, 2)
	query = strings.Replace(query, "@end", end, 2)
	query = strings.Replace(query, "@id", id, 2)
	err = r.db.Raw(query).Scan(&resp).Error

	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (r merchantRepository) FetchMerchantOmzetPagination(id, start, end, offset, limit string) (resp []dto.MerchantOmzetRepoResponse, totalData int64, err error) {
	var baseQuery, query string
	baseQuery = `SELECT day::date, (SELECT merchant_name FROM merchants M WHERE id='@id') as merchant_name, (case when omzet is null then 0 else omzet end)
		FROM(
			SELECT day::date
				FROM generate_series
				('@start'::date
				,'@end'::date
				,'1 day'::interval) day
				) as d
			LEFT JOIN(SELECT T.created_at::date as day, sum(T.bill_total) as omzet
				FROM transactions AS T
				WHERE T.merchant_id = '@id'
				GROUP BY day) o USING (day) offset '@off' limit '@lim' ;`

	query = strings.Replace(baseQuery, "@start", start, 2)
	query = strings.Replace(query, "@end", end, 2)
	query = strings.Replace(query, "@id", id, 2)
	query = strings.Replace(query, "@off", offset, 2)
	query = strings.Replace(query, "@lim", limit, 2)
	err = r.db.Raw(query).Scan(&resp).Error

	if err != nil {
		return resp, 0, err
	}
	totalData = r.db.Raw(query).Scan(&resp).RowsAffected
	return resp, totalData, nil
}
