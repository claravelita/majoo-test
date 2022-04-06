package external

import (
	"github.com/claravelita/majoo-test/domain"
	"gorm.io/gorm"
	"sync"
	"time"
)

func SeedData(db *gorm.DB) {
	userSeed := UserSeedData()

	wg := sync.WaitGroup{}
	var totalUser int64

	wg.Add(1)
	go func() {
		db.Model(&domain.User{}).Count(&totalUser)
		defer wg.Done()
	}()

	if totalUser != int64(len(userSeed)) {
		db.Create(&userSeed)
	}
}

// ID 1 Pass : 12345678
// ID 2 Pass : 87654321
func UserSeedData() []domain.User {
	return []domain.User{
		{
			ID:         1,
			Name:       "Clara Velita Pranolo",
			UserName:   "claravelita",
			Password:   "$2a$10$dSnFxsSJnDsuYlwCrn62cOiD33OfC4P5eCUIRVRmpOq9LiPNtLbvq",
			AuditTable: domain.AuditTable{},
			Merchants: domain.Merchant{
				ID:           1,
				UserID:       1,
				MerchantName: "Merchalbum",
				AuditTable:   domain.AuditTable{},
				Outlets: []domain.Outlet{
					{
						ID:         1,
						MerchantID: 1,
						OutletName: "Merchalbum Bekasi",
						AuditTable: domain.AuditTable{},
						Transactions: []domain.Transaction{
							{
								ID:         1,
								MerchantID: 1,
								OutletID:   1,
								BillTotal:  200000,
								AuditTable: domain.AuditTable{
									CreatedAt: time.Now().AddDate(0, -5, -2),
									CreatedBy: 0,
									UpdatedAt: time.Time{},
									UpdatedBy: 0,
								},
							},
							{
								ID:         2,
								MerchantID: 1,
								OutletID:   1,
								BillTotal:  150000,
								AuditTable: domain.AuditTable{
									CreatedAt: time.Now().AddDate(0, -5, -2),
									CreatedBy: 0,
									UpdatedAt: time.Time{},
									UpdatedBy: 0,
								},
							},
							{
								ID:         3,
								MerchantID: 1,
								OutletID:   1,
								BillTotal:  400000,
								AuditTable: domain.AuditTable{
									CreatedAt: time.Now().AddDate(0, -5, 0),
									CreatedBy: 0,
									UpdatedAt: time.Time{},
									UpdatedBy: 0,
								},
							},
							{
								ID:         4,
								MerchantID: 1,
								OutletID:   1,
								BillTotal:  50000,
								AuditTable: domain.AuditTable{
									CreatedAt: time.Now().AddDate(0, -5, 0),
									CreatedBy: 0,
									UpdatedAt: time.Time{},
									UpdatedBy: 0,
								},
							},
							{
								ID:         5,
								MerchantID: 1,
								OutletID:   1,
								BillTotal:  140000,
								AuditTable: domain.AuditTable{
									CreatedAt: time.Now().AddDate(0, -5, 5),
									CreatedBy: 0,
									UpdatedAt: time.Time{},
									UpdatedBy: 0,
								},
							},
							{
								ID:         6,
								MerchantID: 1,
								OutletID:   1,
								BillTotal:  80000,
								AuditTable: domain.AuditTable{
									CreatedAt: time.Now().AddDate(0, -5, 8),
									CreatedBy: 0,
									UpdatedAt: time.Time{},
									UpdatedBy: 0,
								},
							},
							{
								ID:         7,
								MerchantID: 1,
								OutletID:   1,
								BillTotal:  990000,
								AuditTable: domain.AuditTable{
									CreatedAt: time.Now().AddDate(0, -5, 10),
									CreatedBy: 0,
									UpdatedAt: time.Time{},
									UpdatedBy: 0,
								},
							},
						},
					},
					{
						ID:         2,
						MerchantID: 1,
						OutletName: "Merchalbum Jakarta",
						AuditTable: domain.AuditTable{},
						Transactions: []domain.Transaction{
							{
								ID:         8,
								MerchantID: 1,
								OutletID:   2,
								BillTotal:  400000,
								AuditTable: domain.AuditTable{
									CreatedAt: time.Now().AddDate(0, -5, -3),
									CreatedBy: 0,
									UpdatedAt: time.Time{},
									UpdatedBy: 0,
								},
							},
							{
								ID:         9,
								MerchantID: 1,
								OutletID:   2,
								BillTotal:  4800000,
								AuditTable: domain.AuditTable{
									CreatedAt: time.Now().AddDate(0, -5, -2),
									CreatedBy: 0,
									UpdatedAt: time.Time{},
									UpdatedBy: 0,
								},
							},
							{
								ID:         10,
								MerchantID: 1,
								OutletID:   2,
								BillTotal:  800000,
								AuditTable: domain.AuditTable{
									CreatedAt: time.Now().AddDate(0, -5, 0),
									CreatedBy: 0,
									UpdatedAt: time.Time{},
									UpdatedBy: 0,
								},
							},
							{
								ID:         11,
								MerchantID: 1,
								OutletID:   2,
								BillTotal:  900000,
								AuditTable: domain.AuditTable{
									CreatedAt: time.Now().AddDate(0, -5, 2),
									CreatedBy: 0,
									UpdatedAt: time.Time{},
									UpdatedBy: 0,
								},
							},
							{
								ID:         12,
								MerchantID: 1,
								OutletID:   2,
								BillTotal:  700000,
								AuditTable: domain.AuditTable{
									CreatedAt: time.Now().AddDate(0, -5, 10),
									CreatedBy: 0,
									UpdatedAt: time.Time{},
									UpdatedBy: 0,
								},
							},
							{
								ID:         13,
								MerchantID: 1,
								OutletID:   2,
								BillTotal:  450000,
								AuditTable: domain.AuditTable{
									CreatedAt: time.Now().AddDate(0, -5, 15),
									CreatedBy: 0,
									UpdatedAt: time.Time{},
									UpdatedBy: 0,
								},
							},
							{
								ID:         14,
								MerchantID: 1,
								OutletID:   2,
								BillTotal:  500000,
								AuditTable: domain.AuditTable{
									CreatedAt: time.Now().AddDate(0, -5, 15),
									CreatedBy: 0,
									UpdatedAt: time.Time{},
									UpdatedBy: 0,
								},
							},
							{
								ID:         15,
								MerchantID: 1,
								OutletID:   2,
								BillTotal:  800000,
								AuditTable: domain.AuditTable{
									CreatedAt: time.Now().AddDate(0, -5, 17),
									CreatedBy: 0,
									UpdatedAt: time.Time{},
									UpdatedBy: 0,
								},
							},
							{
								ID:         16,
								MerchantID: 1,
								OutletID:   2,
								BillTotal:  670000,
								AuditTable: domain.AuditTable{
									CreatedAt: time.Now().AddDate(0, -5, 19),
									CreatedBy: 0,
									UpdatedAt: time.Time{},
									UpdatedBy: 0,
								},
							},
							{
								ID:         17,
								MerchantID: 1,
								OutletID:   2,
								BillTotal:  1020000,
								AuditTable: domain.AuditTable{
									CreatedAt: time.Now().AddDate(0, -5, 20),
									CreatedBy: 0,
									UpdatedAt: time.Time{},
									UpdatedBy: 0,
								},
							},
						},
					},
				},
			},
		},
		{
			ID:         2,
			Name:       "Irene",
			UserName:   "renebae",
			Password:   "$2a$10$USEl.o9CytW2KXw5jCxPr.dK5O/G7MVtsqPnvfGNdcgXv9s8FZMIq",
			AuditTable: domain.AuditTable{},
			Merchants: domain.Merchant{
				ID:           2,
				UserID:       2,
				MerchantName: "Goodies",
				AuditTable:   domain.AuditTable{},
				Outlets: []domain.Outlet{
					{
						ID:         3,
						MerchantID: 2,
						OutletName: "Goodies Tangerang",
						AuditTable: domain.AuditTable{},
						Transactions: []domain.Transaction{
							{
								ID:         18,
								MerchantID: 2,
								OutletID:   3,
								BillTotal:  40000,
								AuditTable: domain.AuditTable{
									CreatedAt: time.Now().AddDate(0, -5, -3),
									CreatedBy: 0,
									UpdatedAt: time.Time{},
									UpdatedBy: 0,
								},
							},
							{
								ID:         19,
								MerchantID: 2,
								OutletID:   3,
								BillTotal:  65000,
								AuditTable: domain.AuditTable{
									CreatedAt: time.Now().AddDate(0, -5, -2),
									CreatedBy: 0,
									UpdatedAt: time.Time{},
									UpdatedBy: 0,
								},
							},
							{
								ID:         20,
								MerchantID: 2,
								OutletID:   3,
								BillTotal:  150000,
								AuditTable: domain.AuditTable{
									CreatedAt: time.Now().AddDate(0, -5, 1),
									CreatedBy: 0,
									UpdatedAt: time.Time{},
									UpdatedBy: 0,
								},
							},
							{
								ID:         21,
								MerchantID: 2,
								OutletID:   3,
								BillTotal:  45000,
								AuditTable: domain.AuditTable{
									CreatedAt: time.Now().AddDate(0, -5, 3),
									CreatedBy: 0,
									UpdatedAt: time.Time{},
									UpdatedBy: 0,
								},
							},
							{
								ID:         22,
								MerchantID: 2,
								OutletID:   3,
								BillTotal:  89000,
								AuditTable: domain.AuditTable{
									CreatedAt: time.Now().AddDate(0, -5, 8),
									CreatedBy: 0,
									UpdatedAt: time.Time{},
									UpdatedBy: 0,
								},
							},
							{
								ID:         23,
								MerchantID: 2,
								OutletID:   3,
								BillTotal:  90000,
								AuditTable: domain.AuditTable{
									CreatedAt: time.Now().AddDate(0, -5, 10),
									CreatedBy: 0,
									UpdatedAt: time.Time{},
									UpdatedBy: 0,
								},
							},
							{
								ID:         24,
								MerchantID: 2,
								OutletID:   3,
								BillTotal:  40000,
								AuditTable: domain.AuditTable{
									CreatedAt: time.Now().AddDate(0, -5, 12),
									CreatedBy: 0,
									UpdatedAt: time.Time{},
									UpdatedBy: 0,
								},
							},
						},
					},
				},
			},
		},
	}
}
