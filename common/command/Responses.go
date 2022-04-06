package command

import (
	"github.com/claravelita/majoo-test/common"
	"github.com/claravelita/majoo-test/common/models"
)

func SuccessResponses(data interface{}) (result models.JSONResponses) {
	return models.JSONResponses{
		Status: "Success",
		Data:   data,
		Code:   common.SucceedCode,
	}
}

func NotFoundResponses(status interface{}) (result models.JSONResponses) {
	return models.JSONResponses{
		Status: status,
		Data:   nil,
		Code:   common.NotFoundCode,
	}
}

func BadRequestResponses(data string) (result models.JSONResponses) {
	return models.JSONResponses{
		Status: data,
		Data:   nil,
		Code:   common.BadRequestCode,
	}
}

func InternalServerResponses(data string) (result models.JSONResponses) {
	return models.JSONResponses{
		Status: data,
		Data:   nil,
		Code:   common.InternalServerCode,
	}
}

func SuccessPaginationResponses(data interface{}, meta interface{}) (result models.JSONResponsesPagination) {
	return models.JSONResponsesPagination{
		Status: "Succeed",
		Code:   common.SucceedCode,
		Data:   data,
		Meta:   meta,
	}
}

func InternalServerPaginationResponses(data string) (result models.JSONResponsesPagination) {
	return models.JSONResponsesPagination{
		Status: data,
		Code:   common.InternalServerCode,
		Data:   nil,
		Meta:   models.Meta{},
	}
}

func NotFoundPaginationResponses(status string, data interface{}) (result models.JSONResponsesPagination) {
	return models.JSONResponsesPagination{
		Status: data,
		Code:   common.NotFoundCode,
		Data:   data,
		Meta:   models.Meta{},
	}
}

func BadRequestPaginationResponses(data string) (result models.JSONResponsesPagination) {
	return models.JSONResponsesPagination{
		Status: data,
		Code:   common.BadRequestCode,
		Data:   nil,
		Meta:   models.Meta{},
	}
}
