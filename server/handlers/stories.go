package handlers

import (
	"net/http"
	"strconv"

	"github.com/asxcandrew/herd/server/services"

	"github.com/asxcandrew/herd/server/models"
	"github.com/gin-gonic/gin"
)

type storyResponse struct {
	HTMLBody string `json:"html_body" binding:"required"`
	Title    string `json:"title" binding:"required"`
}

func CreateStory(c *gin.Context) {
	var json storyResponse

	if err := c.BindJSON(&json); err == nil {
		userID, _ := strconv.ParseUint(c.Keys["userID"].(string), 10, 32)

		story := &models.Story{
			AuthorID: userID,
			Title:    json.Title,
			HTMLBody: json.HTMLBody,
		}

		err := services.CreateStory(story)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusBadRequest, "data": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": http.StatusCreated, "data": story})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "data": err})
	}
}

func UpdateStory(c *gin.Context) {
	var json storyResponse
	story := &models.Story{}
	findStory(c, story)

	if err := c.BindJSON(&json); err == nil {
		if json.Title != "" {
			story.Title = json.Title
		}
		if json.HTMLBody != "" {
			story.HTMLBody = json.HTMLBody
		}

		err := models.Update(story)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "data": err})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "data": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusCreated, "data": story})
}

func DeleteStory(c *gin.Context) {
	story := &models.Story{}
	findStory(c, story)

	userID, _ := strconv.ParseUint(c.Keys["userID"].(string), 10, 32)

	if story.AuthorID != userID {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "data": ""})
		return
	}

	err := models.DB().Delete(story)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "data": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": ""})
}

func GetStory(c *gin.Context) {
	story := &models.Story{}
	findStory(c, story)

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
	story := &models.Story{}
	findStory(c, story)

	res := gin.H{
		"id":        c.Param("id"),
		"html_body": story.HTMLBody,
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": res})
}

func findStory(c *gin.Context, story *models.Story) {
	storyID, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": ""})
		return
	}

	story.ID = storyID
	err = models.DB().Select(story)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": ""})
		return
	}
}
