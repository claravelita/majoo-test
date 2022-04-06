package api

import (
	"github.com/claravelita/majoo-test/api/controller"
	"github.com/claravelita/majoo-test/api/handler"
	customMiddleware "github.com/claravelita/majoo-test/api/middleware"
	"github.com/claravelita/majoo-test/common/helper"
	"github.com/claravelita/majoo-test/infrastructure/external"
	"github.com/claravelita/majoo-test/repository/postgres"
	"github.com/claravelita/majoo-test/usecase/merchant"
	"github.com/claravelita/majoo-test/usecase/outlet"
	"github.com/claravelita/majoo-test/usecase/user"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"time"
)

// Server Struct
type Server struct {
	Route *echo.Echo
}

// NewServer : Create Server Instance
func NewServer(e *echo.Echo) *Server {
	return &Server{
		e,
	}
}

func (server *Server) InitializeServer() {
	server.Route.Use(customMiddleware.UseCorsMiddleware())
	customMiddleware.UseCustomLogger(server.Route)
	handler.UseCustomValidatorHandler(server.Route)

	newDB := external.NewGormDB()
	apiGroup := server.Route.Group("/api")

	UserRepo := postgres.NewUserRepository(newDB)
	MerchantRepo := postgres.NewMerchantRepository(newDB)
	OutletRepo := postgres.NewOutletRepository(newDB)

	AuthHelper, _ := helper.NewAuthHelper(os.Getenv("SECRET_AUTH"))

	UserUsecase := user.NewUserImplementation(UserRepo, AuthHelper)
	UserController := controller.NewUserController(UserUsecase)
	UserController.Route(apiGroup)

	MerchantUsecase := merchant.NewMerchantImplementation(MerchantRepo)
	MerchantController := controller.NewMerchantController(MerchantUsecase)
	MerchantController.Route(apiGroup)

	OutletUsecase := outlet.NewOutletImplementation(OutletRepo, MerchantRepo)
	OutletController := controller.NewOutletController(OutletUsecase)
	OutletController.Route(apiGroup)

	serverConfiguration := &http.Server{
		Addr:         ":" + os.Getenv("SERVER_PORT"),
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	server.Route.Logger.Fatal(server.Route.StartServer(serverConfiguration))

}
