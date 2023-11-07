package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Add value to debtors
// @Security ApiKeyAuth
// @Tags transaction
// @Id transaction
// @Accept json
// @Produce json
// @Param uuid path string true "child user"
// @Param value path float32 true "value debtors"
// @Success 200
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/user/{uuid}/transaction{value} [post]
func (h *Handler) addTransaction(c *gin.Context) {
	childUuid := c.Param("uuid")
	if childUuid == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "uuid is null")
	}
	valueStr := c.Param("value")
	if valueStr == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "uuid is null")
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}
	userUUID, err := getUserUUID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}
	err = h.service.User.AddTransaction(c, userUUID, childUuid, value)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}

	c.AbortWithStatus(http.StatusOK)
}
