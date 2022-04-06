package external

import (
	"fmt"
	"github.com/claravelita/majoo-test/common"
	"github.com/claravelita/majoo-test/common/command"
	"github.com/claravelita/majoo-test/common/exception"
	"github.com/claravelita/majoo-test/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func NewGormDB() *gorm.DB {
	dsn := command.PostgreConfig()

	if dbInstance == nil {
		var err error

		DBInstance, err := gorm.Open(postgres.New(dsn), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: false,
		})
		if err != nil {
			exception.PanicIfNeeded(err)
		}

		if common.IsMigrate && dbInstance == nil {
			err = DBInstance.AutoMigrate(&domain.User{}, &domain.Merchant{}, &domain.Outlet{}, &domain.Transaction{})
			exception.PanicIfNeeded(err)
			fmt.Println(err)

			if err == nil {
				fmt.Println("Migrate Suceed!")
			}
		}

		dbInstance = DBInstance
		SeedData(dbInstance)
		return dbInstance
	} else {
		return dbInstance
	}
}
