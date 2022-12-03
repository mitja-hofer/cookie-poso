package routes

import (
	"CookiePoso/controllers"
	"github.com/gin-gonic/gin"
)

func PublicRoutes(g *gin.RouterGroup) {
	g.GET("/login", controllers.LoginGetHandler())
	g.POST("/login", controllers.LoginPostHandler())
	g.GET("/", controllers.IndexGetHandler())
	g.GET("/signup", controllers.SignUpGetHandler())
	g.POST("/signup", controllers.SignUpPostHandler())
}

func PrivateRoutes(g *gin.RouterGroup) {
	g.GET("/dashboard", controllers.DashboardGetHandler())
	g.GET("/logout", controllers.LogoutGetHandler())
	g.GET("/new-recipe", controllers.NewRecipeGetHandler())
	g.POST("/new-recipe", controllers.NewRecipePostHandler())
}
