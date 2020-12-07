package main // temporary main

import (
	"fmt"
	"sync"
	"time"

	"github.com/tkobil/earnings_report/internal"
)

func gatherSecurityInfo(securities []internal.Security, ch chan int) {
	var wg sync.WaitGroup

	for secIdx := range securities {
		//runtime.Breakpoint()
		wg.Add(1)
		go internal.FetchPolygon(&securities[secIdx], secIdx, ch, &wg)
		time.Sleep(12 * time.Second)
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
			//runtime.Breakpoint()
			str := securities[secIdx].SplitByLengthThreshold(300)
			fmt.Println(str)
		case false:
			fmt.Println("No Mo")
			return
		}

	}
}
