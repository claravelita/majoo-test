package domain

type Merchant struct {
	ID           int64  `json:"id" gorm:"AUTO_INCREMENT;NOT NULL;PRIMARY_KEY"`
	UserID       int64  `json:"user_id" gorm:"NOT NULL"`
	MerchantName string `json:"merchant_name" gorm:"NOT NULL;type:varchar(40);"`

	AuditTable   AuditTable    `json:"-" gorm:"embedded"`
	Outlets      []Outlet      `json:"outlets" gorm:"foreignkey:MerchantID;references:ID"`
	Transactions []Transaction `json:"transactions"  gorm:"foreignkey:MerchantID;references:ID"`
}
