package main

import "time"

func ResetGoroutine() {
	for {
		SetNextResetTime(time.Now().Add(ResetInterval))

		time.Sleep(ResetInterval)

		DeleteAllIPCounts()
	}
}
