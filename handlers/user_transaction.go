package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Pin child user UUID to auth user
// @Security ApiKeyAuth
// @Tags user
// @Id pin-user
// @Accept json
// @Produce json
// @Param uuid path string true "child user"
// @Success 200
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/user/pin{uuid} [post]
func (h *Handler) pinUserToUser(c *gin.Context) {
	id := c.Param("uuid")
	if id == "" {
		newErrorResponse(c, http.StatusBadRequest, "uuid is null")
		return
	}

	userUUID, err := getUserUUID(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.service.User.PinUserToUser(c, userUUID, id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

// @Summary Pin child user UUID to auth user
// @Security ApiKeyAuth
// @Tags user
// @Id value-debtor
// @Accept json
// @Produce json
// @Param uuid path string true "child user"
// @Success 200
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/user/debtor/{uuid} [post]
func (h *Handler) getSumTransactionDebtorUser(c *gin.Context) {
	id := c.Param("uuid")
	if id == "" {
		newErrorResponse(c, http.StatusBadRequest, "uuid is null")
		return
	}

	userUUID, err := getUserUUID(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	value, err := h.service.User.GetSumTransactionDebtor(c, userUUID, id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, value)
}
