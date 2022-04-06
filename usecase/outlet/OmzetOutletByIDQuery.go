package outlet

import (
	"github.com/claravelita/majoo-test/common"
	"github.com/claravelita/majoo-test/common/command"
	"github.com/claravelita/majoo-test/common/dto"
	"github.com/claravelita/majoo-test/common/exception"
	"github.com/claravelita/majoo-test/common/models"
	"math"
	"strconv"
	"time"
)

func (u outletImplementation) OmzetOutletByIDQuery(userID string, id string, request dto.DatePaginationParameter) (models.JSONResponsesPagination, error) {
	_, err := time.Parse(common.DateFormat, request.StartDate)
	if err != nil {
		return command.BadRequestPaginationResponses("Start date parameter format should be YYYY-MM-DD"), nil
	}
	_, err = time.Parse(common.DateFormat, request.EndDate)
	if err != nil {
		return command.BadRequestPaginationResponses("End date parameter format should be YYYY-MM-DD"), nil
	}

	if request.PageNumber == 0 {
		request.PageNumber = common.DefaultPageNumber
	}
	if request.PageSize == 0 {
		request.PageSize = common.DefaultPageSize
	}

	offset := request.PageSize * (request.PageNumber - 1)

	user, _ := strconv.Atoi(userID)
	outlet, _ := strconv.Atoi(id)
	findOutlet, err := u.repo.FetchOutletByIDAndUser(int64(user), int64(outlet))
	if err != nil {
		return command.InternalServerPaginationResponses("Internal Server Error"), err
	}
	if findOutlet == nil {
		return command.NotFoundPaginationResponses("User didnt has outlet", nil), nil
	}

	if findOutlet.ID != int64(outlet) {
		return exception.ForbiddenExceptionPagination(), nil
	}

	merchantID := strconv.FormatInt(findOutlet.MerchantID, 10)

	omzet, totalData, err := u.repo.FetchOutletOmzetByID(merchantID, id, request.StartDate, request.EndDate, strconv.Itoa(offset), strconv.Itoa(request.PageSize))
	if err != nil {
		return command.InternalServerPaginationResponses("Internal Server Error"), err
	}

	c := float64(totalData) / float64(request.PageSize)
	totalPages := int(math.Ceil(c))
	if totalPages == 0 {
		totalPages = 1
	}

	conditionNextPage := true
	if request.PageNumber >= totalPages {
		conditionNextPage = false
	}

	meta := models.Meta{
		PageSize:        request.PageSize,
		TotalPages:      totalPages,
		TotalItem:       int(totalData),
		HasNextPage:     conditionNextPage,
		HasPreviousPage: !(request.PageNumber <= 1),
	}

	return command.SuccessPaginationResponses(omzet, meta), nil
}
