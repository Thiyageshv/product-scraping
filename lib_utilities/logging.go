package utilities

import (
	"io"
	"os"
)

func InitializeLogFile(logfilepath string) (*os.File, error) {
	LogFile, err := os.OpenFile(logfilepath,
		os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return nil, err
	}
	io.MultiWriter(os.Stdout, LogFile)
	return LogFile, nil
}

func GetHostIP() string {
	return os.Getenv("HOSTIP")
}


func GetScraperServiceName() string {
	return os.Getenv("SCRAPEHOST")
}