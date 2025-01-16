package rateAlgos


import (
	"time"
	"sync"
)

type FixedWindowCounter struct {
	WindowDefault 			time.Duration
	LimitDefault 			int
	currentRequestCount		int		
	LastResetTime 			time.Time
	mut						sync.Mutex
}

var WindowInfo *FixedWindowCounter


func (fwc *FixedWindowCounter) HandleRequests() bool {
	fwc.resetCounter()
	if WindowInfo.currentRequestCount < WindowInfo.LimitDefault {
		WindowInfo.currentRequestCount += 1
		return true
	}
	return false
}

func (fwc *FixedWindowCounter) resetCounter() {
	if time.Since(WindowInfo.LastResetTime) >= WindowInfo.WindowDefault {
		WindowInfo.LastResetTime = time.Now()
		WindowInfo.currentRequestCount = 0
	}
}
