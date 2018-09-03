package handlers

import (
	"net/http"
	"strconv"

	"github.com/asxcandrew/herd/server/models"
	"github.com/asxcandrew/herd/server/services"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type createMediaStruct struct {
	ContentType string `json:"content_type"binding:"required"`
}

func CreateMedia(c *gin.Context) {
	var json createMediaStruct

	if err := c.BindJSON(&json); err == nil {
		objectName := uuid.NewV4()
		presignedURL, err := services.Storage().PresignedPutObject(objectName.String())

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": ""})
		}

		media := &models.Media{
			Name:        objectName.String(),
			ContentType: json.ContentType,
		}

		models.Create(media)

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": media, "upload_url": presignedURL.String()})
	} else {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": err})
	}
}

func GetMedia(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	media := &models.Media{ID: uint64(id)}
	err := models.DB().Select(media)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": err})
	} else {
		resp := &models.MediaResponse{
			ID:          media.ID,
			ContentType: media.ContentType,
			URL:         services.Storage().GetUrlForObject(media.Name).String(),
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": resp})
	}
}

func DeleteMedia(c *gin.Context) {

}
