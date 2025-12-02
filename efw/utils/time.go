package utils

import "time"

func BeginningOfToday() time.Time {
	t := time.Now()
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func StartAndEndOfDayUTC(subDays ...int) (time.Time, time.Time) {
	now := time.Now().UTC()
	if len(subDays) > 0 {
		now = now.AddDate(0, 0, -subDays[0])
	}
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	end := start.Add(24*time.Hour - time.Nanosecond)
	return start, end
}
