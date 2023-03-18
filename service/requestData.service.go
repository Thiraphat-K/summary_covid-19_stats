package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Thiraphat-K/summary_covid-19_stats/model"
)

func requestData(url string) model.Data {
	method := "GET"
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Printf("Error created HTTP Request: %s", err)
		return model.Data{}
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error requested to URL: %s", err)
		return model.Data{}
	}
	defer res.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	body := buf.Bytes()

	var data model.Data
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("Error convert encoded(JSON byte slice): %s", err)
		return model.Data{}
	}
	
	return data
}
