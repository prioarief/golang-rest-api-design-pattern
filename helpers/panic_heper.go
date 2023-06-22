package helpers

import "github.com/gin-gonic/gin"

func PanicHelper(message string, c *gin.Context) {
	Response(c, 0, message, nil, 500)
}
