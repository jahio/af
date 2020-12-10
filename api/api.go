package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

// Init - set up the API and routes with handlers (and middleware)
func Init(appDB *gorm.DB) http.Handler {
	db = appDB

	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(authenticateUser)

	// Declare routes
	router.GET("/ping", healthCheckHandler)

	return router
}

func authenticateUser(c *gin.Context) {
	// TODO
	c.Next()
}

func healthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": 200, "message": "all systems go"})
}
