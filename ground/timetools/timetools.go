package timetools

import (
	"time"
)

func Str2Time(input string) time.Time {
	layout := "2006-01-02 15:04:05Z07:00"
	output, _ := time.Parse(layout, input)
	return output
}

func Time2Str(input time.Time) string {
	layout := "2006-01-02 15:04:05Z07:00"
	return input.Format(layout)
}
