package outlet

import (
	"github.com/claravelita/majoo-test/common"
	"github.com/claravelita/majoo-test/common/command"
	"github.com/claravelita/majoo-test/common/dto"
	"github.com/claravelita/majoo-test/common/models"
	"strconv"
	"time"
)

func (u outletImplementation) OmzetOutletQuery(userID string, request dto.AnalyticsDateParameter) (models.JSONResponses, error) {
	_, err := time.Parse(common.DateFormat, request.StartDate)
	if err != nil {
		return command.BadRequestResponses("Start date parameter format should be YYYY-MM-DD"), nil
	}
	_, err = time.Parse(common.DateFormat, request.EndDate)
	if err != nil {
		return command.BadRequestResponses("End date parameter format should be YYYY-MM-DD"), nil
	}

	id, _ := strconv.Atoi(userID)
	merchant, err := u.repoMerchant.FetchMerchantByUser(int64(id))
	if err != nil {
		return command.InternalServerResponses("Internal Server Error"), err
	}

	var resp dto.OutletAnalyticsResponse
	resp.MerchantName = merchant.MerchantName
	var outlets []dto.OutletOmzetResponse
	for _, i := range merchant.Outlets {
		var outlet dto.OutletOmzetResponse
		outlet.OutletName = i.OutletName
		omzets, err := u.repo.FetchOutletOmzet(strconv.FormatInt(i.ID, 10), request.StartDate, request.EndDate)
		if err != nil {
			return command.InternalServerResponses("Internal Server Error"), err
		}
		var days []dto.DayOmzetResponse
		for _, o := range omzets {
			var day dto.DayOmzetResponse
			day.Day = o.Day
			day.Omzet = o.Omzet
			days = append(days, day)
		}
		outlet.Omzets = days
		outlets = append(outlets, outlet)
	}
	resp.Outlets = outlets

	return command.SuccessResponses(resp), nil
}
