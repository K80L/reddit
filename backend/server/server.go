package server

import (
	"net/http"

	"github.com/K80L/reddit/backend/api"
	"github.com/K80L/reddit/backend/store"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	r.RedirectTrailingSlash = true

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin,access-control-allow-headers,Authorization"},
	}))

	r.Use(CustomErrors)
	r.POST("/signup", gin.Bind(store.User{}), api.SignUp)
	r.POST("/login", gin.Bind(store.User{}), api.Login)

	router := r.Group("/api", Protect)

	router.GET("/post", api.GetPosts)
	router.POST("/post", gin.Bind(store.Post{}), api.CreatePost)
	router.POST("/post/:post_id/like", api.LikePost)
	router.POST("/post/:post_id/dislike", api.DislikePost)

	router.POST("/subreddit", gin.Bind(store.Subreddit{}), api.CreateSubreddit)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
