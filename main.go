package main

import (
	"log"
	"net/http"

	"github.com/Thiraphat-K/summary_covid-19_stats/controller"
	"github.com/Thiraphat-K/summary_covid-19_stats/service"
	"github.com/gin-gonic/gin"
)

var (
	summaryService    service.SummaryService       = service.NewSummaryService()
	summaryController controller.SummaryController = controller.NewSummaryController(summaryService)
)

func main() {
	routes := gin.Default()

	routes.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"name":    "Donut",
			"message": "SaWasDeeKrub",
		})
	})

	routes.GET("/covid/summary", summaryController.SummaryStats)

	err := routes.Run("localhost:2023")
	if err != nil {
		log.Fatalf("Something went wrong: %s", err)
	}
}
