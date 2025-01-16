package config

import (
	"gopkg.in/yaml.v3"
	"os"
)


type Config struct {

	TokenBucketConfig struct {
		Capacity			int `yaml:"capacity"`
		RefilRate			int `yaml:"refill_rate"`
	} `yaml:"token_bucket"`

	FixedWindowCounterConfig struct {
		WindowDuration		int `yaml:"window_duration"`
		MaxRequest			int `yaml:"max_requests"`
	} `yaml:"fixed_window_counter"`

	SlidingWindowLogConfig struct {
		WindowDuration		int `yaml:"window_duration"`
		MaxRequest			int `yaml:"max_requests"`
	} `yaml:"sliding_window_log"`

}


var config Config

func NewConfig() *Config{
	return &config
}

func Value() *Config{
	return &config 
}

func (c *Config) LoadConfigValues(){
	file , err := os.Open("config.yaml")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)

	if err != nil {
		panic(err)
	}
}

