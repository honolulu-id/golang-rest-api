package controllers

import (
	"github.com/labstack/echo/v4"
	"golang-rest-api/models"
	"net/http"
)

func PariwisataList(c echo.Context) error {

	result, err := models.PariwisataData()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}
