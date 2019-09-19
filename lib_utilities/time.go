package utilities

import (
	"time"
)


func GetCurrentTime() string {
	time_seen := time.Now().UTC().String()
	return time_seen
}