package api

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/K80L/reddit/backend/store"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	user := c.MustGet(gin.BindKey).(*store.User)

	if err := store.AddUser(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Could not sign up"})
		log.Println(err)
		return
	} else {
		setCookie(c, user)
	}
}

func Login(c *gin.Context) {
	user := c.MustGet(gin.BindKey).(*store.User)

	if user, err := store.Authenticate(user.Username, user.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Sign in failed"})
		log.Println(err)
		return
	} else {
		setCookie(c, user)
		c.JSON(http.StatusOK, gin.H{"msg": "Logged in"})
	}
}

func GetUserById(c *gin.Context) {
	userId := c.Param("id")
	id, _ := strconv.Atoi(userId)
	user, _ := store.GetUserById(id)

	c.JSON(http.StatusOK, user)
}

func setCookie(c *gin.Context, user *store.User) {
	token := store.CreateJWT(user)

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    token,
		HttpOnly: true,
		Domain:   "localhost",
		Expires:  time.Now().Add(time.Hour * 24),
	})
}
