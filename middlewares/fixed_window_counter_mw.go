package middlewares

import(
	"fmt"
	"net/http"
	token "github.com/pknrj/RateLimiter/rateAlgos"
	config "github.com/pknrj/RateLimiter/config"
	"time"
)

func FixedWindowCounterMiddleWare(nextHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter , r *http.Request){
			fmt.Println("....Inside Fixed Window Counter middleware....")
			if token.WindowInfo == nil {
				token.WindowInfo = &token.FixedWindowCounter {
					WindowDefault : time.Second * time.Duration(config.Value().FixedWindowCounterConfig.WindowDuration),
					LimitDefault : config.Value().FixedWindowCounterConfig.MaxRequest,
					LastResetTime : time.Now(),
				}
			}

			if ok := token.WindowInfo.HandleRequests() ; ok{
				nextHandler(w , r)
			}else{
				fmt.Fprintf(w, "..Rate Limit Exceeded..\n")
			}
	}
}