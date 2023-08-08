package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/grahamquan/go-crud/models"
	"github.com/grahamquan/go-crud/setup"
)

func PostsCreate(c *gin.Context) {
	// data
	// Title and Body should be uppercase
	// because uppercase means "pubilc"
	var body struct {
		Title string
		Body  string
	}

	err := c.Bind(&body)
	if err != nil {
		fmt.Println(err)
		c.Status(400)
		return
	}
	fmt.Printf("body %v", body)

	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}
	// DB
	result := setup.DB.Create(&post)

	if result.Error != nil {
		fmt.Println(result.Error)
		c.Status(400)
		return
	}
	// response
	c.JSON(200, gin.H{
		"create_count": result.RowsAffected,
	})
}

func PostsGetAll(c *gin.Context) {
	// data
	var posts []models.Post

	// DB
	result := setup.DB.Find(&posts)
	if result.Error != nil {
		fmt.Println(result.Error)
		c.Status(400)
		return
	}

	// response
	c.JSON(200, gin.H{
		"posts": posts,
		"rows":  result.RowsAffected,
	})
}

func PostsGetOne(c *gin.Context) {
	// data
	postId := c.Param("postId")

	// DB
	var post models.Post
	result := setup.DB.First(&post, postId)
	if result.Error != nil {
		fmt.Println(result.Error)
		c.Status(400)
		return
	}

	// response
	c.JSON(200, gin.H{
		"post": post,
		"rows": result.RowsAffected,
	})
}

func PostsUpdate(c *gin.Context) {
	// data
	postId := c.Param("postId")
	var body struct {
		Title string
		Body  string
	}
	err := c.Bind(&body)
	if err != nil {
		fmt.Println(err)
		c.Status(400)
		return
	}

	// DB
	var post models.Post
	// "?" is placeholder for ID
	result := setup.DB.Model(&post).Where("ID = ?", postId).Updates(
		models.Post{
			Title: body.Title,
			Body:  body.Body,
		},
	)

	if result.Error != nil {
		fmt.Println(result.Error)
		c.Status(400)
		return
	}

	// response
	c.JSON(200, gin.H{
		"post": post,
		"rows": result.RowsAffected,
	})
}

func PostsDelete(c *gin.Context) {
	// data
	postId := c.Param("postId")

	// DB
	var post = models.Post{}
	result := setup.DB.Where("ID = ?", postId).Delete(&post)
	if result.Error != nil {
		fmt.Println(result.Error)
		c.Status(400)
		return
	}

	// response
	c.JSON(200, gin.H{
		"delete_count": result.RowsAffected,
	})
}
