package controller

import (
	"net/http"

	"github.com/Thiraphat-K/summary_covid-19_stats/service"
	"github.com/gin-gonic/gin"
)

type SummaryController interface {
	SummaryStats(ctx *gin.Context)
}

type summaryController struct {
	summaryService service.SummaryService
}

func NewSummaryController(summaryService service.SummaryService) SummaryController {
	return &summaryController{
		summaryService: summaryService,
	}
}

func (controller *summaryController) SummaryStats(ctx *gin.Context) {
	data := controller.summaryService.SummaryStats()
	ctx.JSON(http.StatusOK, data)
}
