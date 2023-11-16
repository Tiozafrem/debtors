package handlers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/tiozafrem/debtors/docs"
	"github.com/tiozafrem/debtors/services"
)

type Handler struct {
	service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{service: service}
}

// Add routes for app
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/signUp", h.signUp)
		auth.POST("/signIn", h.signIn)
		auth.POST("/refresh", h.refreshToken)
	}

	api := router.Group("/api", h.userIdentity)
	{
		users := api.Group("/user")
		{
			users.POST("/telegram:id", h.pinTelegramId)
			users.POST("/pin:uuid", h.pinUserToUser)
			users.POST("/debtor/:uuid", h.getSumTransactionDebtorUser)
			users.POST(":uuid/transaction:value", h.addTransaction)
		}
	}
	return router
}
