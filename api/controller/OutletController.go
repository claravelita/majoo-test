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

type outletController struct {
	usecase usecase.OutletUsecase
}

func NewOutletController(usecase usecase.OutletUsecase) *outletController {
	return &outletController{usecase: usecase}
}

func (c *outletController) Route(group *echo.Group) {
	group.GET("/outlet/omzet", c.FetchOmzet, middleware.KeyAuth(authLibs.CheckJWTAndCheckAuthCustomer))
	group.GET("/outlet/omzet/:id", c.FetchOmzetByID, middleware.KeyAuth(authLibs.CheckJWTAndCheckAuthCustomer))
}

func (c *outletController) FetchOmzet(ctx echo.Context) error {
	request := dto.AnalyticsDateParameter{}

	err := models.ValidateRequest(ctx, &request)
	if err != nil {
		return err
	}

	responses, err := c.usecase.OmzetOutletQuery(fmt.Sprint(ctx.Get("user_id")), request)

	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

func (c *outletController) FetchOmzetByID(ctx echo.Context) error {
	request := dto.DatePaginationParameter{}

	err := models.ValidateRequest(ctx, &request)
	if err != nil {
		return err
	}

	parameter := ctx.Param("id")

	responses, err := c.usecase.OmzetOutletByIDQuery(fmt.Sprint(ctx.Get("user_id")), parameter, request)

	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}
