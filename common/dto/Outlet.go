package dto

type (
	OutletOmzetRepoResponse struct {
		Day          string  `json:"day"`
		MerchantName string  `json:"merchant_name"`
		OutletName   string  `json:"outlet_name"`
		Omzet        float64 `json:"omzet"`
	}

	OutletOmzetResponse struct {
		OutletName string             `json:"outlet_name"`
		Omzets     []DayOmzetResponse `json:"omzets"`
	}

	OutletAnalyticsResponse struct {
		MerchantName string                `json:"merchant_name"`
		Outlets      []OutletOmzetResponse `json:"outlets"`
	}
)
