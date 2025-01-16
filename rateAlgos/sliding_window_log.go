package rateAlgos

import (
	"time"
	"sync"
	"log"
)

type SlidingWindowLog struct {
	WindowSize		time.Duration
	MaxRequests 	int
	timeStamps		[]time.Time
	mut 			sync.Mutex
}


var Swl *SlidingWindowLog


func (s *SlidingWindowLog) HandleRequests() bool {
	log.Println("Request Received ")

	now := time.Now()
	s.timeStamps = append(s.timeStamps , now)
	
	s.cleanUp(now)

	if len(s.timeStamps) <= s.MaxRequests {
		return true
	}
	return false
}

func (s *SlidingWindowLog) cleanUp(t time.Time){
	
	s.mut.Lock()
	defer s.mut.Unlock()
	startTime := t.Add(-s.WindowSize)

	var result []time.Time

	for _,val := range(s.timeStamps) {
		if val.After(startTime) {
			result = append(result , val)
		}
	}
	s.timeStamps = make([]time.Time , len(result))
	copy(s.timeStamps , result)
}