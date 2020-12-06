package internal // temporary main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//var basecalendar string = "https://finance.yahoo.com/calendar/earnings/" //Generates today's calendar
var basecalendar string = "https://finance.yahoo.com/calendar/earnings?from=2020-11-29&to=2020-12-05&day=2020-11-30" //If you need to override day

// Security contains metadata about a company's listing
type Security struct {
	Ticker                string
	companyname           string
	latestReportTimestamp string
	latestReportTitle     string
	latestReportURL       string
	latestReportSource    string
}

func getJSONResponse(response string) string {
	var responseslice []string
	responselist := strings.Split(response, "\n")

	for _, value := range responselist {
		if strings.HasPrefix(value, "root.App.main = ") {
			responseslice = append(responseslice, value)
		}
	}
	jsonResp := strings.Split(responseslice[0][:len(responseslice[0])-1], "root.App.main = ")
	return string(jsonResp[1])
}

// GetTodaysReporters Gathers slice of security structs based on todays reporters
func GetTodaysReporters() []Security {
	var securities []Security // initialize return
	response, err := http.Get(basecalendar)
	if err != nil {
		fmt.Println(err) // To-Do: change to logging
		return securities
	}

	defer response.Body.Close()
	var result map[string]interface{}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err) // To-Do: change to logging
		return securities
	}

	responseString := string(responseData)
	getJSONResponse(responseString)
	json.Unmarshal([]byte(getJSONResponse(responseString)), &result)
	rows := result["context"].(map[string]interface{})["dispatcher"].(map[string]interface{})["stores"].(map[string]interface{})["ScreenerResultsStore"].(map[string]interface{})["results"].(map[string]interface{})["rows"] //deep get from JSON
	rowslist := rows.([]interface{})

	for _, value := range rowslist {

		Ticker := value.(map[string]interface{})["ticker"]
		companyname := value.(map[string]interface{})["companyshortname"]
		newsecurity := Security{Ticker: fmt.Sprint(Ticker), companyname: fmt.Sprint(companyname)}
		securities = append(securities, newsecurity)

	}
	return securities

}
