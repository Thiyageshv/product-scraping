package app 


type Config struct {
	LogFilePath       string  	`yaml:"logfilepath"`
	Datetimeformat    string  	`yaml:"datetimeformat"`
	Utctimeformat     string  	`yaml:"utctimeformat"`
	ScrapeInterval 	  int64   	`yaml:"scrapeinterval"`
	ScrapeEndpoint 	  string 	`yaml:"scrapeendpoint"`
	ExpiryThreshold	  float64   `yaml:"expirythreshold"`
	RetryInterval 	  int64  	`yaml:"retryinterval"`
}