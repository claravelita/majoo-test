package api

import (
	"github.com/claravelita/majoo-test/api/handler"
	customMiddleware "github.com/claravelita/majoo-test/api/middleware"
	"github.com/claravelita/majoo-test/infrastructure/external"
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

	external.NewGormDB()
	//apiGroup := server.Route.Group("/api")

	serverConfiguration := &http.Server{
		Addr:         ":" + os.Getenv("SERVER_PORT"),
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	server.Route.Logger.Fatal(server.Route.StartServer(serverConfiguration))

}
