package main

import (
	"errors"
	"time"
)

func unixNanoToMilliseconds(nanoseconds int64) int64 {
	return nanoseconds / (int64(time.Millisecond) / int64(time.Nanosecond))
}

func unixMillisecondsToTime(milliseconds int64) time.Time {
	return time.Unix(0, milliseconds*int64(time.Millisecond))
}

func parseTime(s string) (time.Time, error) {
	patterns := []string{
		time.RFC3339,
		"2006-01-02T15:04:05",
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
		time.RFC822,
		time.RFC850,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		"02 January 2006",
		"January 02 2006",
		"02 Jan 2006",
		"Jan 02 2006",
		"01-02-2006",
		"2006-01-02 15:04:05.999999999 -0700 MST",
		"2006-01-02",
		"2006-01-02T15:04:05-0700",
		"2006-01-02 15:04:05 -07:00",
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05Z07:00",
		"2006-01-02 15:04:05Z0700",
		"2006-01-02 15:04:05",
		"02 Jan 2006 15:04:05 GMT",
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	}
	for _, p := range patterns {
		t, err := time.Parse(p, s)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, errors.New("No matching pattern found for time parsing")
}
