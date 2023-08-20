package controllers

import (
	"crud-go/handlers/rest"
	"crud-go/initializers"
	"crud-go/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func PostsCreate(c echo.Context) error {
	// Get data off req body
	var body models.Post

	err := c.Bind(&body)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Failed to read body",
		})
	}

	if err = body.Validate(); err != nil {
		return rest.ValidationError(c, err)
	}

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)
	// Return it

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Failed to create post",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"post": post,
	})
}

func PostsIndex(c echo.Context) error {
	// Get the posts
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "User not found",
		})
	}

	// Respond with them
	return c.JSON(http.StatusOK, echo.Map{
		"posts": posts,
	})
}

func PostsShow(c echo.Context) error {

	id := c.Param("id")

	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Usuário não encontrado",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"post": post,
	})
}

func PostsUpdate(c echo.Context) error {
	
	id := c.Param("id")

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	var post models.Post
	initializers.DB.First(&post, id)

	err := initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Erro ao atualizarpost",
		})
	}

	return c.JSON(200, echo.Map{
		"post": post,
	})
}

func PostsDelete(c echo.Context) error {

	id := c.Param("id")


	err := initializers.DB.Delete(&models.Post{}, id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Erro ao deletar post",
		})
	}

	return c.JSON(200, "")
}
