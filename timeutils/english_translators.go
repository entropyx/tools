package timeutils

import (
	"fmt"
	"time"
)

func MonthToEnglish(month time.Month) string {
	return fmt.Sprintf("%s", month)
}

func WeekdayToEnglish(weekday time.Weekday) string {
	return fmt.Sprintf("%s", weekday)
}
