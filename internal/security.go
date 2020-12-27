package internal

import (
	"fmt"
	"strings"
)

// Security contains metadata about a company's listing
type Security struct {
	Ticker                string
	companyname           string
	latestReportTimestamp string
	latestReportTitle     string
	latestURLInfo         string
	latestReportSource    string
	_message              string //private message attr
}

func (security *Security) setMessage() {
	message := fmt.Sprintf("%s (%s) Reporting Earnings Today", security.companyname, security.Ticker)
	if len(security.latestReportTitle) > 0 {
		message += fmt.Sprintf(
			"\n\nLatest News (source: %s): %s", security.latestReportSource, security.latestReportTitle)
	}
	security._message = message
}

//This is really just a property - Todo: explore clever way to assign property "message"
func (security *Security) getMessage() string {
	if len(security._message) > 0 {
		return security._message
	}
	security.setMessage()
	return security._message
}

func (security *Security) isAboveLengthThreshold(threshold int) bool {
	return len(security.getMessage())+len(security.latestURLInfo) > threshold
}

//SplitByLengthThreshold will split string message into a slice of threshold-length strings
func (security *Security) SplitByLengthThreshold(threshold int) []string {
	if !security.isAboveLengthThreshold(threshold) {
		return []string{security.getMessage()}
	}

	var tweets []string
	tweetIdx := 0
	words := strings.Split(security._message, " ")
	for _, word := range words {
		switch true {
		case len(tweets) == 0:
			tweets = append(tweets, word)
		case len(tweets[tweetIdx]+" "+word) > (threshold - 7):
			//Time for New Tweet
			tweets = append(tweets, word)
			tweetIdx++
		default:
			tweets[tweetIdx] += " " + word
		}
	}

	// Add Url
	if len(tweets[tweetIdx])+2+len(security.latestURLInfo) > (threshold - 7) { //2 for newlines
		tweets[tweetIdx] += "\n\n" + security.latestURLInfo
	} else {
		tweets = append(tweets, security.latestURLInfo)
		tweetIdx++
	}

	//Add Tweet Header "(1/3): " Now that we know how many tweets we are sending
	for idx := range tweets {
		tweets[idx] = fmt.Sprintf("(%v/%v): ", idx+1, len(tweets)) + tweets[idx]
	}
	return tweets
}
