// The server package is responsible for initializing the Gin server, setting up routes, and applying middleware.
// This is the entry point for our backend application server where we define the API endpoints and attach handlers.
package server

import (
	"log"
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
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Content-Type,access-control-allow-origin,access-control-allow-headers,Authorization"},
		AllowCredentials: true,
	}))

	db := store.GetConnection()

	userStore, err := store.NewUserStore(db)
	if err != nil {
		log.Fatal("Error initializing user store")
	}

	subredditStore, err := store.NewSubredditStore(db)
	if err != nil {
		log.Fatal("Error initializing subreddit store")
	}

	postStore, err := store.NewPostStore(db)
	if err != nil {
		log.Fatal("Error initializing post store")
	}

	postHandler := api.NewPostHandler(postStore)
	userHandler := api.NewUserHandler(userStore)
	subredditHandler := api.NewSubredditHandler(subredditStore)

	// public APIs
	r.Use(CustomErrors)
	// gin.Bind(store.User{}) binds the request body to a User struct and store in the context.
	r.POST("/signup", gin.Bind(store.User{}), userHandler.SignUp)
	r.POST("/login", gin.Bind(store.User{}), userHandler.Login)

	// protected APIs
	protected := r.Group("/api", Protect(userStore))

	// User
	protected.GET("/user/:id", userHandler.GetUserById)
	protected.GET("/user/authenticate", userHandler.CheckIfLoggedIn)
	protected.GET("/user/logout", userHandler.Logout)

	// Post
	protected.GET("/post", postHandler.GetPosts)
	// gin.Bind(store.Post{}) is a Middleware that tells Gin to automatically bind (parse and validate)
	// the request body into a Post struct and make it available in the context.
	protected.POST("/post", gin.Bind(store.Post{}), postHandler.CreatePost)
	// protected.POST("/post/:id/like", postHandler.LikePost)
	// protected.POST("/post/:id/dislike", postHandler.DislikePost)

	// Subreddit
	protected.POST("/subreddit", gin.Bind(store.Subreddit{}), subredditHandler.CreateSubreddit)
	protected.GET("/subreddit/:id", subredditHandler.GetSubreddit)

	// Health check
	protected.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
