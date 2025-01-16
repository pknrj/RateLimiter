package middlewares

import(
	"fmt"
	"net/http"
	"log"
	token "github.com/pknrj/RateLimiter/rateAlgos"
	config "github.com/pknrj/RateLimiter/config"
	"time"
)


func SlidingWindowLogMiddleWare(nextHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter , r *http.Request){
			log.Println("....Inside Sliding Window Log Middleware....")
			
			if token.Swl == nil {
				token.Swl = &token.SlidingWindowLog {
					WindowSize : time.Second * time.Duration(config.Value().SlidingWindowLogConfig.WindowDuration),
					MaxRequests : config.Value().SlidingWindowLogConfig.MaxRequest, 
				}
			}

			if ok := token.Swl.HandleRequests(); ok {
				nextHandler(w , r)
			}else{
				fmt.Fprintf(w, "....rate limit exceeded...\n")
			}
	}
}
