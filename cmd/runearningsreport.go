package main // temporary main

import (
	"fmt"
	"sync"

	"github.com/tkobil/earnings_report/internal"
)

func gatherSecurityInfo(securities []internal.Security, ch chan int) {
	var wg sync.WaitGroup

	for secIdx, security := range securities {
		go internal.FetchPolygon(&security, secIdx, ch, &wg)
		wg.Add(1)
	}

	wg.Wait()
}

func main() {
	ch := make(chan int)
	securities := internal.GetTodaysReporters()
	gatherSecurityInfo(securities, ch)

	for {
		switch secIdx, ok := <-ch; ok {
		case true:
			//make tweet
			fmt.Println(securities[secIdx])
			continue
		case false:
			break
		}

	}
}
