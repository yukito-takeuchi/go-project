package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yourusername/go-project/handlers"
)

// SetupRoutes configures all application routes
func SetupRoutes(e *echo.Echo) {
	// Health check endpoint
	e.GET("/health", handlers.HealthCheck)

	// API v1 group
	v1 := e.Group("/api/v1")
	{
		// Add your API routes here
		// Example:
		// v1.GET("/users", handlers.GetUsers)
		// v1.POST("/users", handlers.CreateUser)
		// v1.GET("/users/:id", handlers.GetUser)
		// v1.PUT("/users/:id", handlers.UpdateUser)
		// v1.DELETE("/users/:id", handlers.DeleteUser)

		// Placeholder route
		v1.GET("/hello", func(c echo.Context) error {
			return c.JSON(200, map[string]string{
				"message": "Hello from Echo API!",
			})
		})
	}
}
