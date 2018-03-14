package timeutils

import "time"

func AddMonth(date time.Time) time.Time {
	year, month, day := date.Date()
	nextMonth := month
	if month < time.December {
		nextMonth = month + 1
	} else {
		nextMonth = time.January
		year++
	}

	location := date.Location()
	firstOfNextMonth := time.Date(year, nextMonth, 1, 0, 0, 0, 0, location)
	lastOfNextMonth := firstOfNextMonth.AddDate(0, 1, -1)

	if day >= lastOfNextMonth.Day() {
		return lastOfNextMonth
	}

	return time.Date(year, nextMonth, day, 0, 0, 0, 0, location)
}
