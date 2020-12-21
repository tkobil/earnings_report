package internal

import (
	"fmt"
	"math"
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
			"\n\nLatest News (source: %s): %s\n\n%s", security.latestReportSource, security.latestReportTitle, security.latestURLInfo)
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
	return len(security.getMessage()) > threshold
}

//SplitByLengthThreshold will split string message into a slice of threshold-length strings
func (security *Security) SplitByLengthThreshold(threshold int) []string {
	if !security.isAboveLengthThreshold(threshold) {
		return []string{security.getMessage()}
	}
	slicelen := int(math.Ceil(float64(len(security.getMessage())) / float64(threshold-7))) //-7 for the "(1/3): " on top of each msg
	stringslice := make([]string, slicelen)
	for i := 0; i < slicelen; i++ {
		strpt := (threshold - 7) * i
		endpt := strpt + threshold - 7
		if endpt > len(security.getMessage()) {
			endpt = len(security.getMessage())
		}
		stringslice[i] = fmt.Sprintf("(%d/%d): ", i+1, slicelen) + security.getMessage()[strpt:endpt]
	}
	return stringslice
}
