package rest

import (
	"crud-go/models"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidationError(c echo.Context, err error) error {
	var errors []models.ValidationError

	for _, e := range err.(validator.ValidationErrors) {
		log.Println(e)

		var element models.ValidationError

		element.Field = e.StructNamespace()
		element.Tag = e.Tag()
		element.Value = e.Param()

		errors = append(errors, element)
	}

	return c.JSON(http.StatusBadRequest, echo.Map{
		"error": errors,
	})
}
