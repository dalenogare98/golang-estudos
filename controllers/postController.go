package controllers

import (
	"crud-go/initializers"
	"crud-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body models.Post

	err := c.Bind(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "errinho to read body",
		})
		return
	}

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)
	// Return it

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create post",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found",
		})
		return
	}

	// Respond with them
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get if off url
	id := c.Param("id")
	// Get the posts
	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Usuário não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get the id of the url
	id := c.Param("id")
	// Get the data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	// Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)
	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})
	// Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get the id of the url
	id := c.Param("id")

	// Delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// Respond with
	c.Status(200)
}
