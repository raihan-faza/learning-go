package api

import (
	"github.com/gin-gonic/gin"
	"github.com/raihan-faza/learning-go/initializers"
	"github.com/raihan-faza/learning-go/models"
)

func init() {
	initializers.ConnectDB()
}

func createPost(c *gin.Context) {
	var post models.Post
	err := c.ShouldBindJSON(&post)
	if err != nil {
		c.JSON(
			400, gin.H{
				"message": err,
			},
		)
		return
	}
	action := initializers.DB.Create(&post)
	if action.Error != nil {
		c.JSON(
			400, gin.H{
				"message": action.Error,
			},
		)
		return
	}
	c.JSON(
		200, gin.H{
			"message": "post created",
		},
	)
	return
}

func updatePost(c *gin.Context, postId int) {
	var post models.Post
	query := initializers.DB.First(&post, postId)
	if query.Error != nil {
		c.JSON(
			400, gin.H{
				"message": query.Error,
			},
		)
		return

	}
	err := c.ShouldBindJSON(&post)
	if err != nil {
		c.JSON(
			400, gin.H{
				"message": err,
			},
		)
		return
	}
	action := initializers.DB.Save(&post)
	if action.Error != nil {
		c.JSON(
			400, gin.H{
				"message": action.Error,
			},
		)
		return
	}
	c.JSON(
		200, gin.H{
			"message": "post updated",
		},
	)
	return
}

func deletePost(c *gin.Context, postId int) {
	var post models.Post
	query := initializers.DB.First(&post, postId)
	if query.Error != nil {
		c.JSON(
			400, gin.H{
				"message": query.Error,
			},
		)
		return
	}
	action := initializers.DB.Delete(&post)
	if action.Error != nil {
		c.JSON(
			400, gin.H{
				"message": action.Error,
			},
		)
		return
	}
	c.JSON(
		200, gin.H{
			"message": "post deleted",
		},
	)
	return
}
