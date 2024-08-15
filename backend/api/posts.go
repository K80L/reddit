package api

import (
	"net/http"
	"strconv"

	"github.com/K80L/reddit/backend/store"
	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	posts, err := store.GetPosts()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func CreatePost(c *gin.Context) {
	post := c.MustGet(gin.BindKey).(*store.Post)
	user := c.MustGet("user").(*store.User)
	post.UserID = user.UserID

	if err := store.CreatePost(post); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Could not create post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Post created"})
}

func LikePost(c *gin.Context) {
	postId := c.Param("id")
	postIdInt, err := strconv.ParseInt(postId, 10, 32)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
	}
	user := c.MustGet("user").(*store.User)

	if err := store.LikePost(int(postIdInt), user.UserID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Could not like post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Post liked"})
}

func DislikePost(c *gin.Context) {
	postId := c.Param("id")
	postIdInt, err := strconv.ParseInt(postId, 10, 32)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
	}
	user := c.MustGet("user").(*store.User)

	if err := store.DislikePost(int(postIdInt), user.UserID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Could not dislike post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Post disliked"})
}
