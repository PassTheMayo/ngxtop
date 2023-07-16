package main

import "time"

func ResetGoroutine() {
	for {
		time.Sleep(time.Minute)

		PruneOldRequests()
	}
}
