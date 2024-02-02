package gin

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

// Struct for json message about error
type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	slog.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
