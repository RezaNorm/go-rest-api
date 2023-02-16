package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rezanorm/go-rest/initializers"
	"github.com/rezanorm/go-rest/models"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)
	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get Post
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get id from url
	id := c.Param("id")

	// Get Post
	var post models.Post
	initializers.DB.Find(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	// Get id from url
	id := c.Param("id")

	// Get body from request
	var body struct {
		Title string
		Body  string
	}
	c.Bind(body)

	// Get Post
	var post models.Post
	initializers.DB.Find(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initializers.DB.Find(&post, id)

	initializers.DB.Delete(post)

	c.Status(200)
}
