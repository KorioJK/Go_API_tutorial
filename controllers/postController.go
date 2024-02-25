package controllers

import (
	"example.com/m/initializers"
	"example.com/m/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	if body.Body == "" || body.Title == "" {
		c.JSON(400, gin.H{"error": "Must provide both Body and Title."})
		return
	}
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error})
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post

	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{"posts": posts})
}

func PostsShow(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{"post": post})
}

func PostsUpdate(c *gin.Context) {

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	if body.Body == "" || body.Title == "" {
		c.JSON(400, gin.H{"error": "Must provide both Body and Title."})
		return
	}

	var post models.Post
	id := c.Param("id")
	initializers.DB.First(&post, id)

	//Update
	initializers.DB.Model(&post).Updates(models.Post{Body: body.Body, Title: body.Title})

	c.JSON(200, gin.H{"post": post})
}

func PostsDelete(c *gin.Context) {

	id := c.Param("id")
	initializers.DB.Delete(&models.Post{}, id)

	c.Status(200)
}
