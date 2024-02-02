package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Struct for parse json username and password
type signInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Struct for input refresh and binding out json
type refreshInput struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// @Summary SignUp
// @Description create account, registry
// @Tags auth
// @ID registry
// @Accept json
// @Produce json
// @Param input body signInput true "account info"
// @Success 200 {integer} models.Tokens "tokens"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/signUp [post]
func (h *Handler) signUp(c *gin.Context) {
	var input signInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	tokens, err := h.service.Authorization.SignUp(c, input.Email, input.Password)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, tokens)
}

// @Summary SignIn
// @Description login
// @Tags auth
// @Id login
// @Accept json
// @Produce json
// @Param input body signInput true "credentials"
// @Success 200 {object} models.Tokens "tokens"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/signIn [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	tokens, err := h.service.Authorization.SignIn(input.Email, input.Password)
	if err != nil {

		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, tokens)
}

// @Summary Refresh
// @Description refresh token
// @Tags auth
// @Id refresh
// @Accept json
// @Produce json
// @Param input body refreshInput true "credentials"
// @Success 200 {object} models.Tokens "tokens"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/refresh [post]
func (h *Handler) refreshToken(c *gin.Context) {
	var input refreshInput

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	tokens, err := h.service.Authorization.RefreshToken(input.RefreshToken)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, tokens)
}
