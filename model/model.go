package model

type CovidCase struct {
	ConfirmDate    string `json:"ConfirmDate"`
	No             int    `json:"No"`
	Age            int    `json:"Age"`
	Gender         string `json:"Gender"`
	GenderEn       string `json:"GenderEn"`
	Nation         string `json:"Nation"`
	Province       string `json:"Province"`
	ProvinceId     int    `json:"ProvinceId"`
	District       string `json:"District"`
	ProvinceEn     string `json:"ProvinceEn"`
	StatQuarantine int    `json:"StatQuarantine"`
}

type Data struct {
	Cases []CovidCase `json:"Data"`
}

type StatsSummary struct {
	SumProvince map[string]int `json:"Province"`
	SumAgeGroup AgeGroupStats  `json:"AgeGroup"`
}

type AgeGroupStats struct {
	Age0to30  int `json:"0-30"`
	Age31to60 int `json:"31-60"`
	Age61plus int `json:"61+"`
	AgeNA     int `json:"N/A"`
}
