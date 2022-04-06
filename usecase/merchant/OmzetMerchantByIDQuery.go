package merchant

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

func (u merchantImplementation) OmzetMerchantByIDQuery(userID string, id string, request dto.DatePaginationParameter) (models.JSONResponsesPagination, error) {
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
	merchant, err := u.repo.FetchMerchantByUser(int64(user))
	if err != nil {
		return command.InternalServerPaginationResponses("Internal Server Error"), err
	}
	if merchant == nil {
		return command.NotFoundPaginationResponses("User didnt has merchant", nil), nil
	}

	merchantID := strconv.FormatInt(merchant.ID, 10)

	if merchantID != id {
		return exception.ForbiddenExceptionPagination(), nil
	}

	omzet, totalData, err := u.repo.FetchMerchantOmzetPagination(merchantID, request.StartDate, request.EndDate, strconv.Itoa(offset), strconv.Itoa(request.PageSize))
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
