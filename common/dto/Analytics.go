package dto

type (
	AnalyticsDateParameter struct {
		StartDate string `query:"start_date"`
		EndDate   string `query:"end_date"`
	}

	DatePaginationParameter struct {
		StartDate  string `query:"start_date"`
		EndDate    string `query:"end_date"`
		PageSize   int    `query:"page_number"`
		PageNumber int    `query:"page_number"`
	}

	DayOmzetResponse struct {
		Day   string  `json:"day"`
		Omzet float64 `json:"omzet"`
	}
)
