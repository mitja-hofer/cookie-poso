package controllers

import (
	"CookiePoso/globals"
	"CookiePoso/helpers"
	"CookiePoso/types"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUpGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.HTML(http.StatusBadRequest, "login.html",
				gin.H{
					"content": "Please logout first",
					"user":    user,
				})
			return
		}
		c.HTML(http.StatusOK, "signup.html", gin.H{
			"content": "",
			"user":    user,
		})
	}
}

func SignUpPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.HTML(http.StatusBadRequest, "login.html", gin.H{
				"content": "Please logout first",
			})
			return
		}

		account := types.Account{
			Username: c.PostForm("username"),
			Password: c.PostForm("password"),
			Email:    c.PostForm("email"),
		}

		if helpers.EmptyUserPassEmail(account.Username, account.Password, account.Email) {
			c.HTML(http.StatusBadRequest, "signup.html", gin.H{
				"content": "Parameters can't be empty",
			})
			return
		}

		_, err := helpers.AddUser(account)
		if err != nil {
			c.HTML(http.StatusUnauthorized, "signup.html", gin.H{
				"content": err,
			})
		}

		c.Redirect(http.StatusMovedPermanently, "/login")
	}
}
