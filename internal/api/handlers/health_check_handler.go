package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Response from the health check endpoint
type HealthResponse struct {
    Status string `json:"status" example:"healthy"`
}

// @Summary Check Health
// @Description check if the web service is healthy
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /health [get]
func HealthCheckHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "healthy",
	})
}
