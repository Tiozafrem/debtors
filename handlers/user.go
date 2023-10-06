package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Pin telegram id to auth user
// @Security ApiKeyAuth
// @Tags user
// @Id pin-telegram
// @Accept json
// @Produce json
// @Param id path string true "telegram Id"
// @Success 200
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/user/telegram{id} [post]
func (h *Handler) pinTelegramId(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "id is null")
	}

	userUUID, err := getUserUUID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
	err = h.service.User.PinTelegramId(c, userUUID, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	c.AbortWithStatus(http.StatusOK)
}
