package timeutils

import "time"

type MonthTranslator func(month time.Month) string

type WeekdayTranslator func(weekday time.Weekday) string
