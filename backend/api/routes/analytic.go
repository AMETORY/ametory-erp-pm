package routes

import (
	"ametory-pm/api/handlers"
	"ametory-pm/api/middlewares"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/gin-gonic/gin"
)

func SetupAnalyticRoutes(r *gin.RouterGroup, erpContext *context.ERPContext) {
	handler := handlers.NewAnalyticHandler(erpContext)
	group := r.Group("/analytic")
	group.Use(middlewares.AuthMiddleware(erpContext, true))
	{
		group.POST("/customer-interaction", handler.CustomerInteractionHandler)
		group.POST("/average-time-reply", handler.AverageTimeReplyHandler)
		group.POST("/hourly-customer-interaction", handler.HourlyCustomerInteractionHandler)
		group.POST("/hourly-average-time-reply", handler.HourlyAverageTimeReplyHandler)
	}

}
