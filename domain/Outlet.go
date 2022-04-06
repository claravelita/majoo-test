package domain

type Outlet struct {
	ID         int64  `json:"id" gorm:"AUTO_INCREMENT;NOT NULL;PRIMARY_KEY"`
	MerchantID int64  `json:"merchant_id" gorm:"NOT NULL;"`
	OutletName string `json:"outlet_name" gorm:"type:varchar(40);NOT NULL;"`

	AuditTable   AuditTable    `json:"-" gorm:"embedded"`
	Transactions []Transaction `json:"transactions"  gorm:"foreignkey:OutletID;references:ID"`
}
