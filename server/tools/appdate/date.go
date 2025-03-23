package appdate

import "time"

func GetStartOfDay() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

func GetEndOfDay() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, int(time.Nanosecond*time.Duration(time.Second-1)), now.Location())
}
