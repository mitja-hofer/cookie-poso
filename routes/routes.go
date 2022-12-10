package routes

import (
	"CookiePoso/controllers"
	"github.com/gin-gonic/gin"
)

func PublicRoutes(g *gin.RouterGroup) {
	g.POST("/login", controllers.Login())
	g.POST("/register", controllers.Register())
}

func PrivateRoutes(g *gin.RouterGroup) {
	g.GET("/user", controllers.CurrentUser())
	g.GET("/user/:username", controllers.GetUserByUsername())
	g.POST("/new-recipe", controllers.NewRecipePostHandler())
	g.GET("/list-recipes", controllers.GetRecipesForLoggedIn())
	g.GET("/recipes/ingredient/:ingredient", controllers.GetRecipesByIngredient())
	g.GET("/list-recipes/:userId", controllers.GetRecipesByUserId())
	g.GET("/list-ingredients/:recipeId", controllers.GetIngredientsByRecipeId())
	g.GET("/ingredients/:name", controllers.GetIngredientsByPartialName())
	g.POST("/ingredients", controllers.PostIngredient())
}
