package controllers

import (
	"go-backend/config"
	"go-backend/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET ALL
// func GetCabangs(c *gin.Context) {
// 	var cabangs []models.UserCabang
// 	if err := config.DB.Find(&cabangs).Error; err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(200, cabangs)
// }

func GetCabangs(c *gin.Context) {
	var results []map[string]interface{}

	err := config.DB.
		Table("t_user_cabang").
		Select(`
			t_user_cabang.cabang_id,
			t_user_cabang.cabang,
			t_user_cabang.alias
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
func GetCabang(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var cabang models.UserCabang
	if err := config.DB.First(&cabang, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Cabang not found"})
		return
	}
	c.JSON(200, cabang)
}

// CREATE
func CreateCabang(c *gin.Context) {
	var cabang models.UserCabang
	if err := c.ShouldBindJSON(&cabang); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&cabang).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, cabang)
}

// UPDATE
func UpdateCabang(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var cabang models.UserCabang
	if err := config.DB.First(&cabang, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Cabang not found"})
		return
	}

	var input models.UserCabang
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Model(&cabang).Updates(input).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, cabang)
}

// DELETE
func DeleteCabang(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var cabang models.UserCabang
	if err := config.DB.First(&cabang, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Cabang not found"})
		return
	}

	if err := config.DB.Delete(&cabang).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Cabang deleted"})
}
