build : 
	go build -o bin/RateLimiter

run : build
	./bin/RateLimiter