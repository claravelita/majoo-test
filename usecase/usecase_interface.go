package usecase

import (
	"github.com/claravelita/majoo-test/common/dto"
	"github.com/claravelita/majoo-test/common/models"
)

type (
	UserUsecase interface {
		CreateUserCommand(request dto.UserRequest) (models.JSONResponses, error)
		LoginUserCommand(request dto.UserLoginRequest) (models.JSONResponses, error)
	}

	MerchantUsecase interface {
		OmzetMerchantQuery(userID string, request dto.AnalyticsDateParameter) (models.JSONResponses, error)
		OmzetMerchantByIDQuery(userID string, id string, request dto.DatePaginationParameter) (models.JSONResponsesPagination, error)
	}

	OutletUsecase interface {
		OmzetOutletQuery(userID string, request dto.AnalyticsDateParameter) (models.JSONResponses, error)
		OmzetOutletByIDQuery(userID string, id string, request dto.DatePaginationParameter) (models.JSONResponsesPagination, error)
	}
)
