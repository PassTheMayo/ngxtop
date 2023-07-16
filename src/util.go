package main

import (
	"sort"
	"sync"
	"time"
)

var (
	ipCounts       map[string]int = make(map[string]int)
	ipCountsMutex  *sync.Mutex    = &sync.Mutex{}
	nextReset      time.Time      = time.Now()
	nextResetMutex *sync.Mutex    = &sync.Mutex{}
)

type IPCount struct {
	IPAddress string
	Count     int
}

func SetNextResetTime(t time.Time) {
	nextResetMutex.Lock()

	defer nextResetMutex.Unlock()

	nextReset = t
}

func DeleteAllIPCounts() {
	ipCountsMutex.Lock()
	ipCounts = make(map[string]int)
	ipCountsMutex.Unlock()
}

func IncrementIPCount(addr string) {
	ipCountsMutex.Lock()

	defer ipCountsMutex.Unlock()

	if v, ok := ipCounts[addr]; ok {
		ipCounts[addr] = v + 1
	} else {
		ipCounts[addr] = 1
	}
}

func GetSortedIPCounts() []IPCount {
	ipCountsMutex.Lock()

	ipCountList := make([]IPCount, 0)

	for ipAddress, count := range ipCounts {
		ipCountList = append(ipCountList, IPCount{
			IPAddress: ipAddress,
			Count:     count,
		})
	}

	ipCountsMutex.Unlock()

	sort.SliceStable(ipCountList, func(i, j int) bool {
		return ipCountList[i].Count > ipCountList[j].Count
	})

	return ipCountList
}
