package main

import (
	"CookiePoso/globals"
	"CookiePoso/middleware"
	"CookiePoso/routes"
	//"github.com/gin-contrib/sessions"
	//"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//router.Use(sessions.Sessions("session", cookie.NewStore(globals.Secret)))

	public := router.Group("/")
	routes.PublicRoutes(public)

	private := router.Group("/")
	private.Use(middleware.JwtAuthMiddleware())
	routes.PrivateRoutes(private)

	router.Run("0.0.0.0:8080")
	defer globals.DB.Close()
}
