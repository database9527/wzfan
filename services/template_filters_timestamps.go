package services

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/flosch/pongo2/v6"
)

// Helper to get time from input value, defaults to now if input is not a valid time.
func getTime(in *pongo2.Value) time.Time {
	// Try to interpret 'in' as a Unix timestamp (integer or string)
	if in.IsInteger() {
		return time.Unix(int64(in.Integer()), 0)
	}
	if in.IsString() {
		ts, err := strconv.ParseInt(in.String(), 10, 64)
		if err == nil {
			return time.Unix(ts, 0)
		}
		// Could add more date string parsing here if needed
	}
	// Default to current time if input is not a recognized timestamp format
	return time.Now()
}

// FilterTimestampYear extracts the year from a timestamp.
// Input can be a Unix timestamp (int/string) or defaults to now().
// Usage: {{ some_timestamp_var|timestamp_year }} or {{ ""|timestamp_year }} for current year
func FilterTimestampYear(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	t := getTime(in)
	return pongo2.AsValue(t.Year()), nil
}

// FilterTimestampMonth extracts the month (1-12) from a timestamp.
func FilterTimestampMonth(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	t := getTime(in)
	return pongo2.AsValue(int(t.Month())), nil
}

// FilterTimestampDay extracts the day of the month (1-31) from a timestamp.
func FilterTimestampDay(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	t := getTime(in)
	return pongo2.AsValue(t.Day()), nil
}

// FilterTimestampHour extracts the hour (0-23) from a timestamp.
func FilterTimestampHour(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	t := getTime(in)
	return pongo2.AsValue(t.Hour()), nil
}

// FilterTimestampMinute extracts the minute (0-59) from a timestamp.
func FilterTimestampMinute(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	t := getTime(in)
	return pongo2.AsValue(t.Minute()), nil
}

// FilterTimestampSecond extracts the second (0-59) from a timestamp.
func FilterTimestampSecond(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	t := getTime(in)
	return pongo2.AsValue(t.Second()), nil
}

// FilterTimestampDatetime formats the timestamp as "YYYY-MM-DD HH:MM:SS".
func FilterTimestampDatetime(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	t := getTime(in)
	return pongo2.AsValue(t.Format("2006-01-02 15:04:05")), nil
}

// FilterTimestampRandom generates a random timestamp within the last hour.
// The input value is ignored.
// Usage: {{ ""|timestamp_random }}
func FilterTimestampRandom(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	now := time.Now()
	// Subtract a random duration between 0 and 3600 seconds (1 hour)
	randomSeconds := rand.Int63n(3600)
	randomTime := now.Add(-time.Duration(randomSeconds) * time.Second)
	// Return as Unix timestamp string, as Pongo2 might handle integers and strings differently
	return pongo2.AsValue(fmt.Sprintf("%d", randomTime.Unix())), nil
}
