package utils

import (
	"time"
)

func ParseDateStringWithEnd(dateStr string, loc *time.Location) (time.Time, error) {
	const layout = "2006-01-02"

	parsedDate, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, err
	}

	return time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(), 23, 59, 59, 0, loc), nil
}

func ParseDateStringWithStart(dateStr string, loc *time.Location) (time.Time, error) {
	const layout = "2006-01-02"

	parsedDate, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, err
	}

	return time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(), 0, 0, 0, 0, loc), nil
}

func NowString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func TodayString() string {
	return time.Now().Format("2006-01-02")
}
