package postgres

import (
	"errors"
	"github.com/claravelita/majoo-test/common/dto"
	"github.com/claravelita/majoo-test/domain"
	"github.com/claravelita/majoo-test/repository"
	"gorm.io/gorm"
	"strings"
)

type outletRepository struct {
	db *gorm.DB
}

func NewOutletRepository(db *gorm.DB) repository.OutletRepository {
	return outletRepository{
		db,
	}
}

func (r outletRepository) FetchOutletByIDAndUser(userID, id int64) (*domain.Outlet, error) {
	var outletEntity domain.Outlet
	err := r.db.Table("outlets").
		Select("outlets.*").
		Joins("LEFT JOIN merchants ON outlets.merchant_id = merchants.id").
		Where("outlets.id = ? AND merchants.user_id = ?", id, userID).First(&outletEntity).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &outletEntity, err
}

func (r outletRepository) FetchOutletOmzet(id, start, end string) (resp []dto.OutletOmzetRepoResponse, err error) {
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
	    WHERE T.outlet_id = '@id'
	    GROUP BY day) t USING (day)`

	query = strings.Replace(baseQuery, "@start", start, 2)
	query = strings.Replace(query, "@end", end, 2)
	query = strings.Replace(query, "@id", id, 2)
	err = r.db.Raw(query).Scan(&resp).Error

	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (r outletRepository) FetchOutletOmzetByID(merchantID, id, start, end, offset, limit string) (resp []dto.OutletOmzetRepoResponse, totalData int64, err error) {
	var baseQuery, query string
	baseQuery = `SELECT day::date, (SELECT merchant_name FROM merchants M WHERE id='@merchantID'), (SELECT outlet_name FROM outlets M WHERE id='@id'), (case when omzet is null then 0 else omzet end)
	FROM(
	SELECT day::date
	    FROM generate_series
	    ('@start'::date
		,'@end'::date
	    ,'1 day'::interval) day
	    ) as d
	LEFT JOIN(SELECT T.created_at::date as day, sum(T.bill_total) as omzet
	    FROM transactions AS T
	    WHERE T.outlet_id = '@id'
	    GROUP BY day) t USING (day) offset '@off' limit '@lim' ;`

	query = strings.Replace(baseQuery, "@start", start, 2)
	query = strings.Replace(query, "@end", end, 2)
	query = strings.Replace(query, "@merchantID", merchantID, 2)
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
