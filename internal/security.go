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
	latestReportURL       string
	latestReportSource    string
	messagecache          string
}

func (security *Security) setMessage() {
	message := fmt.Sprintf("%s (%s) Reporting Earnings Today", security.companyname, security.Ticker)
	if len(security.latestReportTitle) > 0 {
		message += fmt.Sprintf(
			"\n\nLatest News (source: %s): %s\nLink: %s", security.latestReportSource, security.latestReportTitle, security.latestReportURL)
	}
	security.messagecache = message
}

func (security *Security) getMessage() string {
	if len(security.messagecache) > 0 {
		return security.messagecache
	}
	security.setMessage()
	return security.messagecache
}

func (security *Security) isAboveLengthThreshold(threshold int) bool {
	return len(security.latestReportTitle) > threshold
}

//SplitByLengthThreshold will split string message into a slice of threshold-length strings
func (security *Security) SplitByLengthThreshold(threshold int) []string {
	if !security.isAboveLengthThreshold(threshold) {
		return []string{security.getMessage()}
	}

	slicelen := int(math.Ceil(float64(len(security.getMessage())/threshold - 5))) //-5 for the (1/3) on top of each msg
	stringslice := make([]string, slicelen)
	for i := 0; i < slicelen; i++ {
		strpt := (threshold - 5) * i
		endpt := strpt + threshold - 5
		stringslice = append(stringslice, fmt.Sprintf("(%d/%d): ", i+1, slicelen)+security.getMessage()[strpt:endpt])
	}
	return stringslice
}
