package controllers

import (
	"CookiePoso/globals"
	"CookiePoso/helpers"
	"CookiePoso/types"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func NewRecipeGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		c.HTML(http.StatusOK, "new-recipe.html", gin.H{
			"content": "Add a new recipe",
			"user":    user,
		})
	}
}

func NewRecipePostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		userId := session.Get(globals.UserId)
		log.Println(user, userId)

		recipe := types.Recipe{
			UserId: userId.(int64),
			Name:   c.PostForm("name"),
			Text:   c.PostForm("text"),
		}
		_, err := helpers.AddRecipe(recipe)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "new-recipe.html", gin.H{
				"content": err,
			})
		}
		c.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
}
