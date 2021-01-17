package helpers

import "time"

// UNIXTimestampFromNow returns the UNIX timestamp in future as per the given minutes
func UNIXTimestampFromNow(minutes int) int64 {
	now := time.Now().Local()
	return now.Add(time.Minute * time.Duration(minutes)).Unix()
}
