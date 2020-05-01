package routes

import (
	"github.com/hysem/charts/config"
	"github.com/hysem/charts/handler"
	"github.com/labstack/echo/v4"
)

// Set sets the routes for the application
func Set(e *echo.Echo) {
	e.GET("/", handler.ShowDashboard)
	e.GET("/cats", handler.GetCatsChartData)
	e.GET("/dogs", handler.GetDogsChartData)

	e.Static("/assets", config.Current.Path.Asset)
}
