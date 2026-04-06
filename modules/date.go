package modules

import "time"

func IsDateBigger(dateStr string) bool {
	layout := "2006-01-02"
	now := time.Now()
	loc := now.Location()

	targetDate, _ := time.ParseInLocation(layout, dateStr, loc)

	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)

	return !targetDate.Before(todayStart)
}
