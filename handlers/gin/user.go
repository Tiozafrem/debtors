package gin

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
		newErrorResponse(c, http.StatusBadRequest, "id is null")
		return
	}

	userUUID, err := getUserUUID(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.service.User.PinTelegramId(c, userUUID, id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

// @Summary Get all users in system
// @Security ApiKeyAuth
// @Tags user
// @Id users
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/users [get]
func (h *Handler) getUsers(c *gin.Context) {

	_, err := getUserUUID(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	users, err := h.service.User.GetUsers(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}
