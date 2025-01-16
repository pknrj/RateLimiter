package apis

import (
	"github.com/gorilla/mux"
    "net/http"
    "log"
    "fmt"
    "encoding/json"
    "github.com/pknrj/RateLimiter/middlewares"
)

type RateServer struct {
	listenAdd		string
}

func NewRateServer(add string) *RateServer{
	return &RateServer{
		listenAdd : add,
	}
}

func (server *RateServer) StartServer(){
	router := server.setRoutes()
	
	log.Println("RateLimiting Server started on port : " , server.listenAdd)
	log.Fatal(http.ListenAndServe(server.listenAdd , router))
}

func (server *RateServer) setRoutes() *mux.Router {
	
	router := mux.NewRouter()

	router.HandleFunc("/list", server.handleList)

	router.HandleFunc("/algo/token_bucket", 
		middlewares.TokenBucketMiddleWare(server.HandleServerRequests))

	router.HandleFunc("/algo/fixed_window_counter", 
		middlewares.FixedWindowCounterMiddleWare(server.HandleServerRequests))

	router.HandleFunc("/algo/sliding_window_log", 
		middlewares.SlidingWindowLogMiddleWare(server.HandleServerRequests))

	return router 
}



func (s *RateServer) handleList(w http.ResponseWriter , r *http.Request) {
	 err := WriteJson(w , []string{"token_bucket" , "fixed_window_counter" , "sliding_window_log"})
	 if err != nil {
	 	fmt.Fprintf(w , "Error writing the response")
	 }
}


func (s *RateServer) HandleServerRequests(w http.ResponseWriter , r *http.Request){
	fmt.Println(".........request is processed..............")
	fmt.Fprintf(w , "Hi")
}

func WriteJson(w http.ResponseWriter , v any) error {
	w.Header().Add("Content-Type" , "application/json")
	w.WriteHeader(200)
	return json.NewEncoder(w).Encode(v)
}