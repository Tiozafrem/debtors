package gin

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
		newErrorResponse(c, http.StatusBadRequest, "uuid is null")
		return
	}
	valueStr := c.Param("value")
	if valueStr == "" {
		newErrorResponse(c, http.StatusBadRequest, "uuid is null")
		return
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userUUID, err := getUserUUID(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.service.User.AddTransaction(c, userUUID, childUuid, value)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
