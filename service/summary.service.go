package service

import (
	"github.com/Thiraphat-K/summary_covid-19_stats/model"
)

type SummaryService interface {
	SummaryStats() model.StatsSummary
}

type summaryService struct{}

func NewSummaryService() SummaryService {
	return &summaryService{}
}

func (summary *summaryService) SummaryStats() model.StatsSummary {
	URL := "http://static.wongnai.com/devinterview/covid-cases.json"
	Data := requestData(URL)
	sumData := SumStats(Data)
	return sumData
}

func SumStats(data model.Data) model.StatsSummary {
	statsProvince := map[string]int{}
	statsAgeGroup := model.AgeGroupStats{}
	for _, v := range data.Cases {
		statsProvince[v.Province] += 1
		if v.Age <= 30 {
			statsAgeGroup.Age0to30 += 1
		} else if v.Age > 30 && v.Age <= 60 {
			statsAgeGroup.Age31to60 += 1
		} else if v.Age > 60 {
			statsAgeGroup.Age61plus += 1
		} else {
			statsAgeGroup.AgeNA += 1
		}
	}
	statsProvince["N/A"] = statsProvince[""]
	delete(statsProvince, "")
	var sum = model.StatsSummary{
		SumProvince: statsProvince,
		SumAgeGroup: statsAgeGroup,
	}
	return sum
}
