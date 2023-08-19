package rest

import (
	"crud-go/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidationError(c *gin.Context, err error) {
	var errors []models.ValidationError

	for _, e := range err.(validator.ValidationErrors) {
		log.Println(e)

		var element models.ValidationError

		element.Field = e.StructNamespace()
		element.Tag = e.Tag()
		element.Value = e.Param()

		errors = append(errors, element)
	}
	
	c.JSON(http.StatusBadRequest, gin.H{
		"error": errors,
	})

	return
}
