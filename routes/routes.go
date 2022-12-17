package routes

import (
	"CookiePoso/controllers"
	"github.com/gin-gonic/gin"
)

func PublicRoutes(g *gin.RouterGroup) {
	g.GET("/", controllers.Index())
	g.POST("/login", controllers.Login())
	g.POST("/register", controllers.Register())
}

func PrivateRoutes(g *gin.RouterGroup) {
	g.GET("/user", controllers.CurrentUser())
	g.GET("/user/name/:username", controllers.GetUserByUsername())
	g.POST("/recipe", controllers.NewRecipePostHandler())
	g.GET("/recipes", controllers.GetRecipesForLoggedIn())
	g.GET("/recipes/ingredient/:ingredient", controllers.GetRecipesByIngredient())
	g.GET("/recipes/name/:name", controllers.GetRecipeByName())
	g.GET("/recipes/name-like/:name", controllers.GetRecipesByPartialName())
	g.GET("/recipes/user-id/:userId", controllers.GetRecipesByUserId())
	g.GET("/ingredients/recipe-id/:recipeId", controllers.GetIngredientsByRecipeId())
	g.GET("/ingredients/name-like/:name", controllers.GetIngredientsByPartialName())
	g.POST("/ingredients", controllers.PostIngredient())
}
