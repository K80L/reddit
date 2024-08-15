package api

import (
	"net/http"

	"github.com/K80L/reddit/backend/store"
	"github.com/gin-gonic/gin"
)

func CreateSubreddit(c *gin.Context) {
	subreddit := c.MustGet(gin.BindKey).(*store.Subreddit)

	if err := store.CreateSubreddit(subreddit); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Could not create subreddit"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Subreddit created"})
}
