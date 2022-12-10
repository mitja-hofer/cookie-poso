package controllers

import (
	"CookiePoso/models"
	"CookiePoso/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := token.ExtractTokenID(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := models.SelectUserByID(userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
	}
}

func GetUserByUsername() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := models.SelectUserByUsername(c.Param("username"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
	}

}
