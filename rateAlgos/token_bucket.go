package rateAlgos


import (
	"time"
	"sync"
	"fmt"
)

type TokenBucket struct {
	Capacity 			int 
	TokenCount			int
	RefilRate 			int
	LastRefillTime		time.Time
	mut 				sync.Mutex
}

var Bucket *TokenBucket

func (tb *TokenBucket) refill(){
	tb.mut.Lock()
	defer tb.mut.Unlock()

	dur := time.Since(tb.LastRefillTime)

	if dur.Seconds() <= 0 {
		return 
	}

	tokens_to_add := (int)(dur.Seconds()) * tb.RefilRate
	tb.TokenCount = min(tb.Capacity , tb.TokenCount + tokens_to_add)

	tb.LastRefillTime = time.Now()
}

func (tb *TokenBucket) HandleRequests() bool{
	tb.refill()
	tb.mut.Lock()
	defer tb.mut.Unlock()

	if tb.TokenCount == 0 {
		fmt.Printf("can not allow the request\n")
		return false
	}
	tb.TokenCount -= 1 

	fmt.Printf("allowing the request\n")
	return true 
}
