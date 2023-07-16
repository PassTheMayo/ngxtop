package main

import (
	"sort"
	"sync"
	"time"
)

var (
	requests      []Request   = make([]Request, 0)
	requestsMutex *sync.Mutex = &sync.Mutex{}
)

type Request struct {
	IPAddress string
	Timestamp time.Time
}

type UniqueRemoteCount struct {
	IPAddress string
	Count     int
}

func AddRequest(addr string) {
	requestsMutex.Lock()

	defer requestsMutex.Unlock()

	requests = append(requests, Request{
		IPAddress: addr,
		Timestamp: time.Now(),
	})
}

func PruneOldRequests() {
	requestsMutex.Lock()

	defer requestsMutex.Unlock()

	for i := 0; i < len(requests); i++ {
		if time.Since(requests[i].Timestamp) < ResetInterval {
			continue
		}

		requests = requests[i:]
	}
}

func GetSortedIPCounts() []UniqueRemoteCount {
	requestsMutex.Lock()

	defer requestsMutex.Unlock()

	ipCountMap := make(map[string]int)

	for _, req := range requests {
		if time.Since(req.Timestamp) > ResetInterval {
			continue
		}

		if v, ok := ipCountMap[req.IPAddress]; ok {
			ipCountMap[req.IPAddress] = v + 1
		} else {
			ipCountMap[req.IPAddress] = 1
		}
	}

	requestList := make([]UniqueRemoteCount, 0)

	for addr, count := range ipCountMap {
		requestList = append(requestList, UniqueRemoteCount{
			IPAddress: addr,
			Count:     count,
		})
	}

	sort.SliceStable(requestList, func(i, j int) bool {
		return requestList[i].Count > requestList[j].Count
	})

	return requestList
}

func GetTotalRequestCount() int {
	requestsMutex.Lock()

	defer requestsMutex.Unlock()

	count := 0

	for _, req := range requests {
		if time.Since(req.Timestamp) > ResetInterval {
			continue
		}

		count++
	}

	return count
}
