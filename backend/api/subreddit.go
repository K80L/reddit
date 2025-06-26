package api

import (
	"net/http"
	"strconv"

	"github.com/K80L/reddit/backend/store"
	"github.com/gin-gonic/gin"
)

type SubredditHandler struct {
	storage store.SubredditStorage
}

func (h *SubredditHandler) CreateSubreddit(c *gin.Context) {
	subreddit := c.MustGet(gin.BindKey).(*store.Subreddit)

	if err := h.storage.CreateSubreddit(subreddit); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Could not create subreddit"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Subreddit created"})
}

func (h *SubredditHandler) GetSubreddit(c *gin.Context) {
	subredditId := c.Param("id")
	id, _ := strconv.Atoi(subredditId)

	subreddit, err := h.storage.GetSubreddit(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Could not get subreddit"})
		return
	}

	c.JSON(http.StatusOK, subreddit)
}

func NewSubredditHandler(storage store.SubredditStorage) *SubredditHandler {
	return &SubredditHandler{storage: storage}
}
