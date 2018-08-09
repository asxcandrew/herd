package handlers

import (
	"net/http"
	"strconv"

	"github.com/asxcandrew/herd/server/services"

	"github.com/asxcandrew/herd/server/models"
	"github.com/gin-gonic/gin"
)

type createStoryResponse struct {
	HTMLBody string `json:"html_body" binding:"required"`
	Title    string `json:"title" binding:"required"`
}

func CreateStory(c *gin.Context) {
	var json createStoryResponse

	if err := c.BindJSON(&json); err == nil {
		userID, _ := strconv.ParseUint(c.Keys["userID"].(string), 10, 32)

		story := &models.Story{
			AuthorID: userID,
			Title:    json.Title,
			HTMLBody: json.HTMLBody,
		}

		err := services.CreateStory(story)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": http.StatusCreated, "data": story})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func UpdateStory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func DeleteStory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func GetStory(c *gin.Context) {
	storyID, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": ""})
		return
	}

	story := &models.Story{ID: storyID}
	err = models.DB().Select(story)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": story})
}

func QueryStories(c *gin.Context) {

	var stories []models.Story
	err := models.QueryStories(&stories)

	if err != nil || len(stories) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": stories})
}

func GetStoryBody(c *gin.Context) {
	storyID, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": ""})
		return
	}

	story := &models.Story{ID: storyID}
	err = models.DB().Model(story).Column("HTMLBody").Select()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": ""})
		return
	}

	res := gin.H{
		"id":        storyID,
		"html_body": story.HTMLBody,
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": res})
}
