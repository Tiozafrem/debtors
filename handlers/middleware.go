package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "UserUUID"
)

func (h *Handler) userIdentity(c *gin.Context) {
	ctx := c.Request.Context()
	token, err := authTokenFromHeader(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err)
	}

	userUUID, err := h.service.ParseTokenToUserUUID(ctx, token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err)
	}

	c.Set(userCtx, userUUID)
	c.Next()
}

func authTokenFromHeader(c *gin.Context) (string, error) {
	headerAuth := c.GetHeader(authorizationHeader)
	if headerAuth == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(headerAuth, " ")
	if len(headerParts) != 2 || strings.ToLower(headerParts[0]) != "bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		newErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return "", errors.New("token is empty")
	}
	return headerParts[1], nil
}
