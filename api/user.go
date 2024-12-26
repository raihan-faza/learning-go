package api

import (
	"github.com/gin-gonic/gin"
	"github.com/raihan-faza/learning-go/initializers"
	"github.com/raihan-faza/learning-go/models"
)

func createUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(
			400, gin.H{
				"message": err,
			},
		)
		return
	}
	action := initializers.DB.Create(&user)
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
			"message": "user created",
		},
	)
	return
}

func updateUser(c *gin.Context, userId int) {
	var user models.User
	query := initializers.DB.First(&user, userId)
	if query.Error != nil {
		c.JSON(
			400, gin.H{
				"message": query.Error,
			},
		)
		return

	}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(
			400, gin.H{
				"message": err,
			},
		)
		return
	}
	action := initializers.DB.Save(&user)
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
			"message": "user updated",
		},
	)
	return
}

func deleteUser(c *gin.Context, userId int) {
	var user models.User
	query := initializers.DB.First(&user, userId)
	if query.Error != nil {
		c.JSON(
			400, gin.H{
				"message": query.Error,
			},
		)
		return
	}
	action := initializers.DB.Delete(&user)
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
			"message": "user deleted",
		},
	)
	return
}
