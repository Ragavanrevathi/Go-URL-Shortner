package handler

import (
	"net/http"
	"shorten-url/internal/models/request"
	"shorten-url/internal/service"
	"shorten-url/internal/utils"
	"shorten-url/pkg/common"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ShortURL(c *gin.Context) {

	var request request.ShortURL

	// Bind and validate the incoming request body
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	common.LogWithTrace(c.Request.Context()).Info("Request Received", zap.Any("body", request))

	// Check If the URL is reachable or Not
	_, err := utils.IsURLReachable(request.URL)
	if err != nil {
		common.LogWithTrace(c.Request.Context()).Warn("Invalid URL", zap.String("url", request.URL))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate the short url
	shortenedURL := service.ShortURL(request.URL)

	c.JSON(http.StatusOK, gin.H{"short_url": shortenedURL})
}

func RedirectURL(c *gin.Context) {
	key := c.Param("shortKey")
	common.LogWithTrace(c.Request.Context()).Info("Request Received for Redirection:", zap.String("Key", key))

	// check if the key is present
	if originalURL, found := service.GetUserURL(key); found {
		c.Redirect(http.StatusFound, originalURL)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
	}
}
