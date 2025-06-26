package api

import (
	"net/http"

	"github.com/K80L/reddit/backend/store"
	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	storage store.PostStorage
}

func NewPostHandler(storage store.PostStorage) *PostHandler {
	return &PostHandler{storage: storage}
}

func (h *PostHandler) GetPosts(c *gin.Context) {
	posts, err := h.storage.GetPosts()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	// Get the bound post struct from the context
	post := c.MustGet(gin.BindKey).(*store.Post)
	user := c.MustGet("user").(*store.User)
	post.UserID = int(user.ID)

	if err := h.storage.CreatePost(post); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Could not create post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Post created"})
}

// func (h *PostHandler) LikePost(c *gin.Context) {
// 	id := c.Param("id")
// 	postID, err := strconv.Atoi(id)

// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
// 		return
// 	}

// 	user := c.MustGet("user").(*store.User)

// 	if user.HasLiked(postID) {
// 		store.UndoLikePost(postID, int(user.ID))
// 		c.JSON(http.StatusOK, gin.H{"msg": "Post unliked"})
// 		return
// 	}

// 	if err := store.LikePost(postID, int(user.ID)); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Could not like post"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"msg": "Post liked"})
// }

// func (h *PostHandler) DislikePost(c *gin.Context) {
// 	id := c.Param("id")
// 	postID, err := strconv.Atoi(id)

// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
// 		return
// 	}

// 	user := c.MustGet("user").(*store.User)

// 	if user.HasDisliked(postID) {
// 		store.UndoDislikePost(postID, int(user.ID))
// 		c.JSON(http.StatusOK, gin.H{"msg": "Post undisliked"})
// 		return
// 	}

// 	if err := store.DislikePost(postID, int(user.ID)); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Could not dislike post"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"msg": "Post disliked"})
// }

// func (h *PostHandler) DeletePost(c *gin.Context) {
// 	id := c.Param("id")
// 	postID, _ := strconv.Atoi(id)

// 	// check if the post belongs to the user
// 	post, err := store.GetPostByID(postID)

// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Could not find post to delete"})
// 		return
// 	}

// 	user := c.MustGet("user").(*store.User)

// 	if post.User.ID != user.ID {
// 		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized to delete this post"})
// 		return
// 	}

// 	if err := store.DeletePost(postID); err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Could not delete post"})
// 		return
// 	}
// }
