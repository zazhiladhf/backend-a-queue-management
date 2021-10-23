package utils

import (
	"strconv"
	"time"
)

var (
	today = time.Now()
)

func FormatGetDate() string {
	tomorrow := today.AddDate(0, 0, 1).Format("02-01-2006")

	return tomorrow
}

func FormatGetToday() string {
	today := today.Format("02-01-2006")

	return today
}

func FormatGetHour() string {
	hr, min, _ := today.Clock()
	hour := strconv.Itoa(hr) + ":" + strconv.Itoa(int(min))

	return hour
}
