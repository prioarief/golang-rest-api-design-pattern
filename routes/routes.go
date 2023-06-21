package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/check-health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	v1 := r.Group("/api/v1")
	{
		TodoRoutes(v1)
	}

	return r
}
