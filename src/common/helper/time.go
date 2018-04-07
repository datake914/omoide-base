package helper

import "time"

func BeginingOfThisMonth(t time.Time) time.Time {
	return t.AddDate(0, 0, -1*t.Day()+1)
}

func BeginingOfNextMonth(t time.Time) time.Time {
	return t.AddDate(0, 1, -1*t.Day()+1)
}

func EndOfLastMonth(t time.Time) time.Time {
	return t.AddDate(0, 0, -1*t.Day())
}

func EndOfThisMonth(t time.Time) time.Time {
	return t.AddDate(0, 1, -1*t.Day())
}
