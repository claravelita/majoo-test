package domain

type User struct {
	ID       int64  `json:"id" gorm:"AUTO_INCREMENT;NOT NULL;PRIMARY_KEY"`
	Name     string `json:"name" gorm:"type:varchar(45);DEFAULT;NULL;"`
	UserName string `json:"user_name" gorm:"type:varchar(45);DEFAULT;NULL;"`
	Password string `json:"password" gorm:"type:varchar(225);DEFAULT;NULL;"`

	AuditTable AuditTable `json:"-" gorm:"embedded"`
	Merchants  Merchant   `json:"merchants" gorm:"foreignkey:UserID;references:ID"`
}
