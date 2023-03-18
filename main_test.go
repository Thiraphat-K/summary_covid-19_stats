package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sort"
	"testing"

	"github.com/Thiraphat-K/summary_covid-19_stats/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

var (
	r = gin.Default()
)

func SetUpRequest(relativePath string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, relativePath, nil)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	return req
}

func TestRequestStatus(t *testing.T) {
	r.GET("/covid/summary", summaryController.SummaryStats)
	req := SetUpRequest("/covid/summary")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusOK)
}

func TestKeysMain(t *testing.T) {
	// r.GET("/covid/summary", summaryController.SummaryStats)
	req := SetUpRequest("/covid/summary")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var testBody = model.StatsSummary{}
	testByte, _ := json.Marshal(testBody)
	var testStruct map[string]interface{}
	json.Unmarshal(testByte, &testStruct)

	buf := new(bytes.Buffer)
	buf.ReadFrom(w.Body)
	respByte := buf.Bytes()
	var respStruct map[string]interface{}
	json.Unmarshal(respByte, &respStruct)

	testVal := reflect.ValueOf(testStruct).MapKeys()
	respVal := reflect.ValueOf(respStruct).MapKeys()
	sort.Slice(testVal, func(i, j int) bool { return testVal[i].String() > testVal[j].String() })
	sort.Slice(respVal, func(i, j int) bool { return respVal[i].String() > respVal[j].String() })
	for i := 0; i < len(respVal); i++ {
		if respVal[i].String() == testVal[i].String() {
			fmt.Printf("%dst Key ResponseBody: %s == %dst Key TestBody: %s (PASS)\n", i+1, respVal[i].String(), i+1, testVal[i].String())
		} else {
			fmt.Printf("%dst Key ResponseBody: %s != %dst Key TestBody: %s (FAIL)\n", i+1, respVal[i].String(), i+1, testVal[i].String())
		}
		assert.Equal(t, respVal[i].String(), testVal[i].String())
	}
}
