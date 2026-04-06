package controllers

import (
	"go-backend/config"
	"go-backend/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET ALL
// func GetGroups(c *gin.Context) {
// 	var groups []models.UserGroup
// 	if err := config.DB.Find(&groups).Error; err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(200, groups)
// }

func GetGroups(c *gin.Context) {
	var results []map[string]interface{}

	err := config.DB.
		Table("t_user_group").
		Select(`
			t_user_group.user_group_id,
			t_user_group.user_group
		`).
		Scan(&results).Error

	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "success",
		"data":   results,
	})
}

// GET BY ID
func GetGroup(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var group models.UserGroup
	if err := config.DB.First(&group, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "group not found"})
		return
	}
	c.JSON(200, group)
}

// CREATE
func CreateGroup(c *gin.Context) {
	var group models.UserGroup
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&group).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, group)
}

// UPDATE
func UpdateGroup(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var group models.UserGroup
	if err := config.DB.First(&group, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "group not found"})
		return
	}

	var input models.UserGroup
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Model(&group).Updates(input).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, group)
}

// DELETE
func DeleteGroup(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var group models.UserGroup
	if err := config.DB.First(&group, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "group not found"})
		return
	}

	if err := config.DB.Delete(&group).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "group deleted"})
}
