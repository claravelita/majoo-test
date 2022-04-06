package domain

type User struct {
	ID       int64  `json:"id" gorm:"AUTO_INCREMENT;NOT NULL;PRIMARY_KEY"`
	Name     string `json:"name" gorm:"DEFAULT NULL;type:varchar(45)"`
	UserName string `json:"user_name" gorm:"DEFAULT NULL;type:varchar(45)"`
	Password string `json:"password" gorm:"DEFAULT NULL;type:varchar(225)"`

	AuditTable AuditTable `json:"-" gorm:"embedded"`
	Merchants  Merchant   `json:"merchants" gorm:"foreignkey:UserID;references:ID"`
}
