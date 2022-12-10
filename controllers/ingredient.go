package controllers

import (
	"CookiePoso/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NewIngredientInput struct {
	Name string `json:"name" binding:"required"`
}

func GetIngredientsByPartialName() gin.HandlerFunc {
	return func(c *gin.Context) {
		ingredients, err := models.SelectIngredientsByPartialName(c.Param("name"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"content": err,
			})
		}

		c.IndentedJSON(http.StatusOK, ingredients)
	}
}

func PostIngredient() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input NewRecipeInput

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		ingredient := models.Ingredient{
			Name: input.Name,
		}

		_, err := models.InsertIngredient(ingredient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"message": "added new ingredient"})
	}
}
