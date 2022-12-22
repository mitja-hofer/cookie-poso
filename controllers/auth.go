package controllers

import (
	"CookiePoso/models"
	"CookiePoso/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input LoginInput

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		u := models.User{}
		u.Username = input.Username
		u.Password = input.Password

		id, ok := models.CheckUserPass(u.Username, u.Password)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad password"})
			return
		}

		token, err := token.GenerateToken(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "username of password is incorrect"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input RegisterInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		user := models.User{}
		user.Username = input.Username
		user.Password = input.Password
		user.Email = input.Email

		_, err := models.AddUser(user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"message": "registration success"})
	}
}

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}
