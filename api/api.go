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
	Title := c.Query("title")
	Description := c.Query("description")
	Detail := c.Query("detail")
	Post := models.Post{
		Title:       Title,
		Description: Description,
		Detail:      Detail,
	}
	result := initializers.DB.Create(&Post)
	if result.Error != nil {
		c.JSON(
			400, gin.H{
				"message": result.Error,
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

func updatePost(c *gin.Context) {
	return
}

func deletePost(c *gin.Context) {
	return
}
