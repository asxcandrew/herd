package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateStory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func UpdateStory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func DeleteStory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func GetStory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func GetRevisionsList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func SetRevision(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
