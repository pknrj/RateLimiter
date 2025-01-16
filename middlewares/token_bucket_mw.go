package middlewares

import(
	"fmt"
	"net/http"
	token "github.com/pknrj/RateLimiter/rateAlgos"
	config "github.com/pknrj/RateLimiter/config"
	"time"

)

func TokenBucketMiddleWare(nextHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter , r *http.Request){
			fmt.Println("Inside token Bucket middleware")
			if token.Bucket == nil {
				token.Bucket = &token.TokenBucket {
					Capacity : config.Value().TokenBucketConfig.Capacity,
					TokenCount : config.Value().TokenBucketConfig.Capacity,
					RefilRate : config.Value().TokenBucketConfig.RefilRate,
					LastRefillTime : time.Now(),
				}
			}
			ok := token.Bucket.HandleRequests()
			if ok{
				fmt.Fprintf(w, "Hello from server !!!!!\n")
				nextHandler(w , r)
			}else{
				fmt.Fprintf(w, "Bucket is empty wait for a sec to make new request !!!!!\n")
			}
	}
}
