package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"error": false,
		"data":  data,
	})
}

func ErrorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"error": true,
		"data":  err.Error(),
	})
}
