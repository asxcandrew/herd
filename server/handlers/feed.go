package handlers

import (
	"net/http"

	"github.com/asxcandrew/herd/server/models"
	"github.com/gin-gonic/gin"
)

func Feed(c *gin.Context) {
	var stories []models.Story
	err := models.QueryStories(&stories).Where("active = ?", true).Order("published_at ASC").Select()

	if err != nil || len(stories) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": stories})
}
