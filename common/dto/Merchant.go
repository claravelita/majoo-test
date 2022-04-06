package dto

type (
	MerchantOmzetRepoResponse struct {
		Day          string  `json:"day"`
		MerchantName string  `json:"merchant_name"`
		Omzet        float64 `json:"omzet"`
	}

	MerchantAnalyticsResponse struct {
		MerchantName string             `json:"merchant_name"`
		Omzets       []DayOmzetResponse `json:"omzets"`
	}
)
