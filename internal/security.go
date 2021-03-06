package internal

import (
	"fmt"
	"math"
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
	_message              string //private body of tweet attr
	_link                 string //private link for tweet attr
}

func (security *Security) setMessage() {
	message := fmt.Sprintf("%s (%s) Reporting Earnings Today", security.companyname, security.Ticker)
	if len(security.latestReportTitle) > 0 {
		message += fmt.Sprintf(
			"\n\nLatest News (source: %s): %s", security.latestReportSource, security.latestReportTitle)
	}

	security._message = message
	security._link = security.latestURLInfo
}

//This is really just a property - Todo: explore clever way to assign property "message"
func (security *Security) getMessage() (string, string) {
	if len(security._message) != 0 && len(security._link) != 0 {
		return security._message, security._link
	}
	security.setMessage()
	return security._message, security._link
}

// This is a property to calculate the lenth of the tweet in number of characters
func (security *Security) getLength() int {
	if len(security._message) == 0 {
		security.setMessage()
	}
	return len(security._message) + len(security._link) + 2 // + 2 is for the two newline chars needed to separate body and link
}

func (security *Security) isAboveLengthThreshold(threshold int) bool {
	return security.getLength() > threshold
}

// SplitByLengthThreshold will split string message into a slice of threshold-length strings
// Ryan: Changed getMessage to return a map. This function now is reponsible for parsing
// the map and formatting the tweets into
func (security *Security) SplitByLengthThreshold(threshold int) []string {
	if !security.isAboveLengthThreshold(threshold) {
		body, link := security.getMessage()

		return []string{body + "\n\n" + link}
	}

	buffer := threshold - 7 // -7 for the "(1/3): " on top of each msg
	tweetIdx := 0

	body, link := security.getMessage()
	words := strings.Split(body, " ")
	numTweets := int(math.Ceil(float64(security.getLength()) / float64(buffer)))
	linklen := len(link)

	tweets := make([]string, numTweets)

	for _, word := range words {
		switch true {
		case len(tweets) == 0:
			tweets[tweetIdx] = fmt.Sprintf("(%d/%d): ", tweetIdx+1, numTweets) + word
		case len(tweets[tweetIdx]+" "+word) > buffer:
			// Time for New Tweet
			tweetIdx++
			tweets[tweetIdx] = fmt.Sprintf("(%d/%d): ", tweetIdx+1, numTweets) + word
		default:
			tweets[tweetIdx] += " " + word
		}
	}

	// Add Url
	if len(tweets[tweetIdx])+2+linklen > (buffer) { // 2 for newlines
		tweets[tweetIdx] += "\n\n" + link
	} else {
		tweetIdx++
		tweets[tweetIdx] = link
	}

	return tweets
}
