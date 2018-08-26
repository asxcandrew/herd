package handlers

import (
	"net/http"

	"github.com/asxcandrew/herd/server/services"

	"github.com/asxcandrew/herd/server/models"
	"github.com/gin-gonic/gin"
)

func CreateStory(c *gin.Context) {
	var json models.StoryResponse

	if err := c.BindJSON(&json); err == nil {
		userID := uint64(c.Keys["userID"].(float64))

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
	var json models.StoryResponse
	story := &models.Story{}
	findStory(c, story)

	userID := uint64(c.Keys["userID"].(float64))

	if story.AuthorID != userID {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	if err := c.BindJSON(&json); err == nil {
		err := services.UpdateStory(story, &json)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "data": err})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "data": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": story})
}

func DeleteStory(c *gin.Context) {
	story := &models.Story{}
	findStory(c, story)

	userID := uint64(c.Keys["userID"].(float64))

	if story.AuthorID != userID {
		c.AbortWithStatus(http.StatusForbidden)
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
	userID := uint64(c.Keys["userID"].(float64))
	err := models.QueryStories(&stories).Where("author_id = ?", userID).Select()

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
		"uid":       story.UID,
		"html_body": story.HTMLBody,
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": res})
}

func findStory(c *gin.Context, story *models.Story) {
	err := models.DB().Model(story).Where("uid = ?", c.Param("uid")).Select()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
