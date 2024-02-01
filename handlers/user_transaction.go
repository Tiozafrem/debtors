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

// @Summary Get sum transaction user
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
// @Router /api/user/debtor/{uuid} [get]
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

// @Summary Get sum transaction users
// @Security ApiKeyAuth
// @Tags user
// @Id value-debtors
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/user/debtors [get]
func (h *Handler) getSumTransactionDebtorsUser(c *gin.Context) {

	userUUID, err := getUserUUID(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	value, err := h.service.User.GetSumTransactionDebtors(c, userUUID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, value)
}

// @Summary Get sum transaction my
// @Security ApiKeyAuth
// @Tags user
// @Id value-my
// @Accept json
// @Produce json
// @Success 200
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/user/my [get]
func (h *Handler) getSumTransactionMy(c *gin.Context) {

	userUUID, err := getUserUUID(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	value, err := h.service.User.GetSumMy(c, userUUID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, value)
}
