package helpers

import (
	"github.com/gin-gonic/gin"
)

type ResponseType struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Response(c *gin.Context, code int, message string, data any, statusCode int) {
	var result ResponseType

	result.Code = code
	result.Message = message
	result.Data = data

	c.JSON(statusCode, result)
}
