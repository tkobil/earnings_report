package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/tkobil/earnings_report/utils"
)

// FetchPolygon fills out security attrs from polygon API
func FetchPolygon(security *Security, secIdx int, ch chan int, wg *sync.WaitGroup) {
	//fmt.Println("FIRST:", security.Ticker)
	polygonAPIKey := os.Getenv("POLYGONAPIKEY")
	if polygonAPIKey == "" {
		log.Fatal("ERROR: you must set the environment variable POLYGONAPIKEY before running this program")
	}

	defer wg.Done()
	polygonHTTPURLSlice := []string{"https://api.polygon.io/v1/meta/symbols/", security.Ticker, "/news?perpage=1&page=1&apiKey=", polygonAPIKey}
	polygonHTTPURL := strings.Join(polygonHTTPURLSlice, "")
	response, err := http.Get(polygonHTTPURL)
	if err != nil {
		utils.Logger.Error(err.Error())
		ch <- secIdx
		return
	}

	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		utils.Logger.Error(err.Error())
		ch <- secIdx
		return
	}

	var result []map[string]interface{}
	responseString := string(responseData)
	json.Unmarshal([]byte(responseString), &result)
	if len(result) <= 0 {
		fmt.Println(security.Ticker)
		utils.Logger.Warning("No Polygon Results Found for " + security.Ticker)
		ch <- secIdx
		return
	}

	latestResult := result[0]
	security.latestReportTimestamp = fmt.Sprint(latestResult["timestamp"])
	security.latestReportTitle = fmt.Sprint(latestResult["title"])
	security.latestURLInfo = fmt.Sprintf("Latest Report: %s", latestResult["url"])
	security.latestReportSource = fmt.Sprint(latestResult["source"])
	ch <- secIdx
	return
}
