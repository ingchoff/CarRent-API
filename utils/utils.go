package utils

import (
	"strconv"
	"time"
)

func ConvertUnixToDateTimeFormat(ts string) string {
	// Unix timestamp in milliseconds
	timestamp, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		panic(err)
	}
	// Convert the timestamp to seconds by dividing by 1000
	seconds := timestamp / 1000

	// Convert the remaining milliseconds to nanoseconds
	nanoseconds := (timestamp % 1000) * int64(time.Millisecond)

	// Create the time.Time object
	dateTime := time.Unix(seconds, nanoseconds)

	// Format the time.Time object to the desired format
	startDate := dateTime.Format("2006-01-02 15:04:05.000")
	return startDate
}