package main // temporary main

import (
	"fmt"
	"sync"
	"time"

	"github.com/tkobil/earnings_report/internal"
)

var apiCallDelay time.Duration = 12

func gatherSecurityInfo(securities []internal.Security, ch chan int) {
	var wg sync.WaitGroup

	for secIdx, _ := range securities {
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
			//make tweet
			fmt.Println(securities[secIdx])
		case false:
			return
		}

	}
}
