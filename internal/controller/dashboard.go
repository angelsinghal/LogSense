package controller

import (
	"net/http"

	"github.com/angelsinghal/LogSense/internal/middleware"
	"github.com/angelsinghal/LogSense/internal/service"
	"github.com/gin-gonic/gin"
)

type DashboardController struct {
	dashboardService *service.DashboardService
}

func NewDashboardController(dashboardService *service.DashboardService) *DashboardController {
	return &DashboardController{dashboardService: dashboardService}
}

func (c *DashboardController) Summary(ctx *gin.Context) {
	summary, err := c.dashboardService.Summary(middleware.TenantID(ctx))
	if handleError(ctx, err) {
		return
	}
	ctx.JSON(http.StatusOK, summary)
}
