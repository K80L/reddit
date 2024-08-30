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
	post.UserID = int(user.ID)

	if err := store.CreatePost(post); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Could not create post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Post created"})
}

func LikePost(c *gin.Context) {
	id := c.Param("id")
	postID, err := strconv.Atoi(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
	}

	user := c.MustGet("user").(*store.User)

	if user.HasLiked(postID) {
		store.UndoLikePost(postID, int(user.ID))
		c.JSON(http.StatusOK, gin.H{"msg": "Post unliked"})
		return
	}

	if err := store.LikePost(postID, int(user.ID)); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Could not like post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Post liked"})
}

func DislikePost(c *gin.Context) {
	id := c.Param("id")
	postID, err := strconv.Atoi(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
	}

	user := c.MustGet("user").(*store.User)

	if user.HasDisliked(postID) {
		store.UndoDislikePost(postID, int(user.ID))
		c.JSON(http.StatusOK, gin.H{"msg": "Post undisliked"})
		return
	}

	if err := store.DislikePost(postID, int(user.ID)); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Could not dislike post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Post disliked"})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	postID, _ := strconv.Atoi(id)

	// check if the post belongs to the user
	post, err := store.GetPostByID(postID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Could not find post to delete"})
	}

	user := c.MustGet("user").(*store.User)

	if post.User.ID != user.ID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized to delete this post"})
	}

	if err := store.DeletePost(postID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Could not delete post"})
	}
}
