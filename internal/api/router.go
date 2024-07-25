package api

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/anuraj2023/bank-account-management-be/internal/api/handlers"
	customMiddleware "github.com/anuraj2023/bank-account-management-be/internal/api/middleware"
	"github.com/anuraj2023/bank-account-management-be/internal/config"
	"github.com/anuraj2023/bank-account-management-be/internal/repository"
)

type Server struct {
	echo *echo.Echo
	cfg  *config.Config
}

func NewServer(cfg *config.Config, repo repository.AccountRepository) *Server {
	e := echo.New()

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(customMiddleware.ErrorHandler)

	// Handler
	accountHandler := handlers.NewAccountHandler(repo)

	// Health Check Route
	e.GET("/health", handlers.HealthCheckHandler)

	// Account Routes
	e.POST("/accounts", accountHandler.CreateAccount)
	e.GET("/accounts/:accountNumber", accountHandler.GetAccount)
	e.GET("/accounts", accountHandler.ListAccounts)

	return &Server{
		echo: e,
		cfg:  cfg,
	}
}

func (s *Server) Start(address string) error {
	return s.echo.Start(address)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}