package api

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/K80L/reddit/backend/store"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	storage store.UserStorage
}

func (h *UserHandler) SignUp(c *gin.Context) {
	user := c.MustGet(gin.BindKey).(*store.User)

	if err := h.storage.AddUser(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Could not sign up"})
		log.Println(err)
		return
	} else {
		setCookie(c, user)
	}
}

func (h *UserHandler) Login(c *gin.Context) {
	user := c.MustGet(gin.BindKey).(*store.User)

	if user, err := h.storage.Authenticate(user.Username, user.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Sign in failed"})
		log.Println(err)
		return
	} else {
		setCookie(c, user)
		c.JSON(http.StatusOK, gin.H{"msg": "Logged in"})
	}
}

func (h *UserHandler) Logout(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
		Domain:   "localhost",
		Expires:  time.Unix(0, 0),
	})

	c.JSON(http.StatusOK, gin.H{"loggedOut": true, "msg": "Logged out"})
}

func (h *UserHandler) GetUserById(c *gin.Context) {
	userId := c.Param("id")
	id, _ := strconv.Atoi(userId)
	user, _ := h.storage.GetUserById(id)

	c.JSON(http.StatusOK, user)
}

func setCookie(c *gin.Context, user *store.User) {
	token := store.CreateJWT(user)

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
		Domain:   "localhost",
		Expires:  time.Now().Add(time.Hour * 24 * 7),
	})
}

func (h *UserHandler) CheckIfLoggedIn(c *gin.Context) {
	username, err := store.ValidateJWT(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not logged in"})
		return
	}

	user, err := h.storage.GetUser(username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"isAuthenticated": true,
		"user":            user,
	})
}

func NewUserHandler(storage store.UserStorage) *UserHandler {
	return &UserHandler{
		storage: storage,
	}
}
