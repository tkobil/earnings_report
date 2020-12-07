package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

var polygonAPIKey string = "apiKey=cqvrOEHpOWFMUKh7iIogOGA9rV53dVmp"

// FetchPolygon fills out security attrs from polygon API
func FetchPolygon(security *Security, secIdx int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	polygonHTTPURLSlice := []string{"https://api.polygon.io/v1/meta/symbols/", security.Ticker, "/news?perpage=50&page=1&", polygonAPIKey}
	polygonHTTPURL := strings.Join(polygonHTTPURLSlice, "")
	response, err := http.Get(polygonHTTPURL)
	if err != nil {
		fmt.Println(err) //To-Do: Change to logging
		ch <- secIdx
		return
	}

	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err) // To-Do: change to logging
		ch <- secIdx
		return
	}

	var result []map[string]interface{}
	responseString := string(responseData)
	json.Unmarshal([]byte(responseString), &result)
	if len(result) <= 0 {
		fmt.Println(security.Ticker)
		fmt.Println("No Results") //Change to logging
		ch <- secIdx
		return
	}
	latestResult := result[0]
	security.latestReportTimestamp = fmt.Sprint(latestResult["timestamp"])
	security.latestReportTitle = fmt.Sprint(latestResult["title"])
	security.latestReportURL = fmt.Sprint(latestResult["url"])
	security.latestReportSource = fmt.Sprint(latestResult["source"])
	ch <- secIdx
	return
}
