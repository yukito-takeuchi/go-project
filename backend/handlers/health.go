package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/yourusername/go-project/models"
)

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Database  string `json:"database"`
}

// HealthCheck handles health check requests
func HealthCheck(c echo.Context) error {
	response := HealthResponse{
		Status:    "ok",
		Timestamp: time.Now().Format(time.RFC3339),
		Database:  "disconnected",
	}

	// Check database connection
	if models.DB != nil {
		sqlDB, err := models.DB.DB()
		if err == nil && sqlDB.Ping() == nil {
			response.Database = "connected"
		}
	}

	return c.JSON(http.StatusOK, response)
}
