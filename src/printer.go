package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bit101/go-ansi"
)

func PrinterGoroutine() {
	for {
		ansi.ClearScreen()
		ansi.MoveTo(0, 0)

		ipCountList := GetSortedIPCounts()

		if len(ipCountList) < 1 {
			fmt.Println("\n There does not seem to be any requests in the last minute.")
		} else {
			fmt.Printf("\n Showing 0-%d of %d unique requests\n\n", 10, len(ipCountList))

			for i, value := range ipCountList {
				if i >= 10 {
					break
				}

				fmt.Printf("  %2d. %s %s %d\n", i+1, value.IPAddress, strings.Repeat("-", (15-len(value.IPAddress)+3)), value.Count)
			}
		}

		nextResetMutex.Lock()

		fmt.Printf("\n Next reset in %s\n", time.Until(nextReset).Round(time.Millisecond))

		nextResetMutex.Unlock()

		time.Sleep(UpdateInterval)
	}
}
