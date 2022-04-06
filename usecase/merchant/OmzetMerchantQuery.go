package merchant

import (
	"github.com/claravelita/majoo-test/common"
	"github.com/claravelita/majoo-test/common/command"
	"github.com/claravelita/majoo-test/common/dto"
	"github.com/claravelita/majoo-test/common/models"
	"strconv"
	"time"
)

func (u merchantImplementation) OmzetMerchantQuery(userID string, request dto.AnalyticsDateParameter) (models.JSONResponses, error) {
	_, err := time.Parse(common.DateFormat, request.StartDate)
	if err != nil {
		return command.BadRequestResponses("Start date parameter format should be YYYY-MM-DD"), nil
	}
	_, err = time.Parse(common.DateFormat, request.EndDate)
	if err != nil {
		return command.BadRequestResponses("End date parameter format should be YYYY-MM-DD"), nil
	}

	id, _ := strconv.Atoi(userID)
	merchant, err := u.repo.FetchMerchantByUser(int64(id))
	if err != nil {
		return command.InternalServerResponses("Internal Server Error"), err
	}

	var resp dto.MerchantAnalyticsResponse
	resp.MerchantName = merchant.MerchantName
	omzets, err := u.repo.FetchMerchantOmzet(strconv.FormatInt(merchant.ID, 10), request.StartDate, request.EndDate)
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
	resp.Omzets = days

	return command.SuccessResponses(resp), nil
}
