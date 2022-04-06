package controller

import (
	"fmt"
	authLibs "github.com/claravelita/majoo-test/api/middleware"
	"github.com/claravelita/majoo-test/common/dto"
	"github.com/claravelita/majoo-test/common/models"
	"github.com/claravelita/majoo-test/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type merchantController struct {
	usecase usecase.MerchantUsecase
}

func NewMerchantController(usecase usecase.MerchantUsecase) *merchantController {
	return &merchantController{usecase: usecase}
}

func (c *merchantController) Route(group *echo.Group) {
	group.GET("/merchant/omzet", c.FetchOmzet, middleware.KeyAuth(authLibs.CheckJWTAndCheckAuthCustomer))
	group.GET("/merchant/omzet/:id", c.FetchOmzetByID, middleware.KeyAuth(authLibs.CheckJWTAndCheckAuthCustomer))
}

func (c *merchantController) FetchOmzet(ctx echo.Context) error {
	request := dto.AnalyticsDateParameter{}

	err := models.ValidateRequest(ctx, &request)
	if err != nil {
		return err
	}

	responses, err := c.usecase.OmzetMerchantQuery(fmt.Sprint(ctx.Get("user_id")), request)

	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

func (c *merchantController) FetchOmzetByID(ctx echo.Context) error {
	request := dto.DatePaginationParameter{}

	err := models.ValidateRequest(ctx, &request)
	if err != nil {
		return err
	}

	parameter := ctx.Param("id")

	responses, err := c.usecase.OmzetMerchantByIDQuery(fmt.Sprint(ctx.Get("user_id")), parameter, request)

	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}
