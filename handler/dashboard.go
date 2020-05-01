package handler

import (
	"net/http"

	"github.com/hysem/charts/model"
	"github.com/hysem/charts/templates"
	"github.com/labstack/echo/v4"
)

// ShowDashboard shows the dashboard
func ShowDashboard(c echo.Context) error {
	response := templates.DefaultParams(c)
	return c.Render(http.StatusOK, "dashboard", response)
}

// GetCatsChartData retrieves the cats data from db
func GetCatsChartData(c echo.Context) error {
	data, err := model.GetCatsChartData()
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, "Unable to complete the request")
	}
	return c.JSON(http.StatusOK, data)
}

// GetDogsChartData retrieves the dogs data from db
func GetDogsChartData(c echo.Context) error {
	data, err := model.GetDogsChartData()
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, "Unable to complete the request")
	}
	return c.JSON(http.StatusOK, data)
}
