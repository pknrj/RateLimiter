package main

import (
	api "github.com/pknrj/RateLimiter/apis"
	file "github.com/pknrj/RateLimiter/config"
)

func main(){
	// load config values from yaml file 
	file.Value().LoadConfigValues()
	
	server := api.NewRateServer(":3000")
	server.StartServer()

}