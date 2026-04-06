package controllers

import (
	"net/http"

	"go-backend/config"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(c *gin.Context) {
	var req LoginRequest

	// bind JSON dari frontend
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Request tidak valid",
		})
		return
	}

	var user struct {
		Username  string
		Password  string
		Name      string
		Email     string
		GroupName string
	}

	// ambil user dari DB
	err := config.DB.
		Table("t_user").
		Select("username, password, name, email, group_name").
		Where("username = ?", req.Username).
		Scan(&user).Error

	if err != nil || user.Username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "User tidak ditemukan",
		})
		return
	}

	// cek password (plain text)
	if req.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Password salah",
		})
		return
	}

	// sukses login
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Login berhasil",
		"user": gin.H{
			"username":   user.Username,
			"name":       user.Name,
			"email":      user.Email,
			"group_name": user.GroupName,
		},
	})
}
