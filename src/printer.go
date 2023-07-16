package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/bit101/go-ansi"
	"golang.org/x/term"
)

func PrinterGoroutine() {
	for {
		ansi.ClearScreen()
		ansi.MoveTo(0, 0)

		fmt.Print("\u001b[?25l")

		w, h, err := term.GetSize(0)

		if err != nil {
			panic(err)
		}

		sortedRequests := GetSortedIPCounts()

		if len(sortedRequests) < 1 {
			fmt.Println("\n There has not been any requests in the last minute.")
		} else {
			totalRequests := GetTotalRequestCount()
			uniqueRequestCount := h - 6

			fmt.Printf("\n Showing 1-%d of %d unique remote hosts\n\n", uniqueRequestCount, len(sortedRequests))

			for i, value := range sortedRequests {
				if i >= uniqueRequestCount {
					break
				}

				percentage := float64(value.Count) / float64(totalRequests)

				fmt.Printf("  %2d. %s %s %4d [%s%s] %.2f%%\n", i+1, value.IPAddress, strings.Repeat(" ", (15-len(value.IPAddress)+3)), value.Count, strings.Repeat("#", int(math.Ceil(percentage*(float64(w)-42)))), strings.Repeat(".", int(math.Floor((1.0-percentage)*(float64(w)-42)))), percentage*100)
			}

			fmt.Printf("%s %d total requests in the last %s", strings.Repeat("\n", int(math.Max(float64(uniqueRequestCount-len(sortedRequests)), 0))+1), totalRequests, time.Duration(math.Min(float64(ResetInterval), float64(time.Since(startedAt)))).Round(time.Second))
		}

		time.Sleep(UpdateInterval)
	}
}
