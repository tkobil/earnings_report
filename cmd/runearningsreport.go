package main // temporary main

import (
	"sync"
	"time"

	"github.com/tkobil/earnings_report/internal"
)

var apiCallDelay time.Duration = 12

func gatherSecurityInfo(securities []internal.Security, ch chan int) {
	var wg sync.WaitGroup

	for secIdx := range securities {
		wg.Add(1)
		go internal.FetchPolygon(&securities[secIdx], secIdx, ch, &wg)
		time.Sleep(apiCallDelay * time.Second)
	}

	wg.Wait()
	close(ch)
}

func main() {
	ch := make(chan int)
	securities := internal.GetTodaysReporters()
	go gatherSecurityInfo(securities, ch)
	for {
		switch secIdx, ok := <-ch; ok {
		case true:
			secstr := securities[secIdx].SplitByLengthThreshold(300) //300 is default max length of tweet
			go internal.SendTweets(secstr)
		case false:
			return
		}

	}
}
