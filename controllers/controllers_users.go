package controllers

import (
	"strconv"

	"go-backend/config"
	"go-backend/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var results []map[string]interface{}

	err := config.DB.
		Table("t_user").
		Select(`
			t_user.username,
			t_user.name,
			t_user.email,
			t_user.group_name,
			t_user_cabang.cabang
		`).
		Joins("LEFT JOIN t_user_cabang ON t_user.cabang_id = t_user_cabang.cabang_id").
		Joins("LEFT JOIN t_user_group ON t_user.group_name = t_user_group.user_group").
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
func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, user)
}

// CREATE
func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, user)
}

// UPDATE
func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&user).Updates(input)

	c.JSON(200, user)
}

// DELETE
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	config.DB.Delete(&user)

	c.JSON(200, gin.H{"message": "User deleted"})
}
