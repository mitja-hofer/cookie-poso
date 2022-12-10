package controllers

import (
	"CookiePoso/models"
	"CookiePoso/token"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type NewRecipeInput struct {
	Name        string                      `json:"name" binding:"required"`
	Text        string                      `json:"text" binding:"required"`
	Ingredients []models.IngredientInRecipe `json:"ingredients" binding:"required"`
}

type IngredientInRecipeInput struct {
	Id     int64  `json:"id"`
	Name   string `json:"name" binding:"required"`
	Amount int64  `json:"amount" binding:"required"`
	Unit   string `json:"unit" binding:"required"`
}

func NewRecipePostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input NewRecipeInput

		userId, err := token.ExtractTokenID(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		recipe := models.Recipe{
			UserId:      userId,
			Name:        input.Name,
			Text:        input.Text,
			Ingredients: input.Ingredients,
		}

		_, err = models.AddRecipe(recipe)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "added new recipe"})
	}
}

func GetRecipesForLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := token.ExtractTokenID(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		recipes, err := models.SelectRecipesByUserId(userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"content": err,
			})
		}

		c.IndentedJSON(http.StatusOK, recipes)
	}
}

func GetRecipesByUserId() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"content": err,
			})
		}
		recipes, err := models.SelectRecipesByUserId(userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"content": err,
			})
		}

		c.IndentedJSON(http.StatusOK, recipes)
	}
}

func GetIngredientsByRecipeId() gin.HandlerFunc {
	return func(c *gin.Context) {
		recipeId, err := strconv.ParseInt(c.Param("recipeId"), 10, 64)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"content": err,
			})
		}
		recipes, err := models.IngredientsByRecipeId(recipeId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"content": err,
			})
		}

		c.IndentedJSON(http.StatusOK, recipes)
	}
}

func GetRecipesByIngredient() gin.HandlerFunc {
	return func(c *gin.Context) {
		recipes, err := models.SelectRecipesByIngredient(c.Param("ingredient"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"content": err,
			})
		}

		c.JSON(http.StatusOK, recipes)
	}
}

func GetRecipeByName() gin.HandlerFunc {
	return func(c *gin.Context) {
		recipe, err := models.SelectRecipeByName(c.Param("name"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"content": err,
			})
		}

		c.JSON(http.StatusOK, recipe)
	}
}

func GetRecipesByPartialName() gin.HandlerFunc {
	return func(c *gin.Context) {
		recipes, err := models.SelectRecipesByPartialName(c.Param("name"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"content": err,
			})
		}

		c.JSON(http.StatusOK, recipes)
	}
}
