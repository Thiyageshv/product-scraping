package app 


type Config struct {
	LogFilePath       string  `yaml:"logfilepath"`
	Datetimeformat    string  `yaml:"datetimeformat"`
	Utctimeformat     string  `yaml:"utctimeformat"`
}