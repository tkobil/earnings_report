package main // temporary main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type security struct {
	ticker      string
	companyname string
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

var basecalendar string = "https://finance.yahoo.com/calendar/earnings/"

func main() {
	response, err := http.Get(basecalendar)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	var result map[string]interface{}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	responseString := string(responseData)
	getJSONResponse(responseString)
	json.Unmarshal([]byte(getJSONResponse(responseString)), &result)

	rows := result["context"].(map[string]interface{})["dispatcher"].(map[string]interface{})["stores"].(map[string]interface{})["ScreenerResultsStore"].(map[string]interface{})["results"].(map[string]interface{})["rows"]
	rowslist := rows.([]interface{})
	var securities []security
	for _, value := range rowslist {

		ticker := value.(map[string]interface{})["ticker"]
		companyname := value.(map[string]interface{})["companyshortname"]
		newsecurity := security{fmt.Sprint(ticker), fmt.Sprint(companyname)}
		securities = append(securities, newsecurity)

	}

	fmt.Println(securities) // ToDo: change func name from name and return this

}
