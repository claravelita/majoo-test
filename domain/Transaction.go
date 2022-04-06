package domain

type Transaction struct {
	ID         int64   `json:"id" gorm:"AUTO_INCREMENT;NOT NULL;PRIMARY_KEY"`
	MerchantID int64   `json:"merchant_id" gorm:"NOT NULL;"`
	OutletID   int64   `json:"outlet_id" gorm:"NOT NULL;"`
	BillTotal  float64 `json:"bill_total" gorm:"NOT NULL;"`

	AuditTable AuditTable `json:"-" gorm:"embedded"`
}
