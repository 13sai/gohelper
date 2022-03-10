package time

import (
	"time"
)

// 2006-01-02 15:04:05 remember the time

const AMAZE_TIME = "2006-01-02 15:04:05"

func DateTimeToTimeStamp(dateTime, format string) int64 {
	t, _ := time.ParseInLocation(format, dateTime, time.Local)
	return t.Unix()
}

func TimeStampToDateTime(ts int64, format string) string {
	return time.Unix(ts, 0).Format(format)
}

func GetMonday(t time.Time) time.Time {
	offset := int(time.Monday - t.Weekday())
	if offset > 0 {
		offset = -6
	}

	year, month, day := t.Date()
	return time.Date(year, month, day+offset, 0, 0, 0, 0, time.Local)
}

func GetFirstDayOfMonth(t time.Time) time.Time {
	year, month, _ := t.Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
}

func GetLastDayOfMonth(t time.Time) time.Time {
	year, month, _ := t.Date()
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.Local)
}

func GetFirstDayOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, time.Local)
}

func GetLastDayOfYear(t time.Time) time.Time {
	return time.Date(t.Year()+1, 1, 0, 0, 0, 0, 0, time.Local)
}

func GetZero(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}
