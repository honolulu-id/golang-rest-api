package controllers

import (
	"github.com/labstack/echo/v4"
	"golang-rest-api/models"
	"net/http"
)

func HeatMapDataList(c echo.Context) error {

	// resultGetGroupByDesanoData, err := models.GetGroupByDesanoData()
	// resultGetPasienMapIcdData, err := models.GetPasienMapIcdData()
	// resultGetPenyakitByKelurahanData, err := models.GetPenyakitByKelurahanData()
	resultGetLatLongDesanoData, err := models.GetLatLongDesanoData()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, resultGetLatLongDesanoData)

}
