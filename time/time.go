package time

import (
	"time"
)

// 2006-01-02 15:04:05 记住这一刻

// 获取周一时间
func GetMonday(t time.Time) time.Time {
	offset := int(time.Monday - t.Weekday())
	if offset > 0 {
		offset = -6
	}
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
}

//获取传入的时间所在月份的第一天
func GetFirstDayOfMonth(t time.Time) time.Time {
	return t.AddDate(0, 0, -t.Day()+1)
}

//获取传入的时间所在月份的最后一天
func GetLastDayOfMonth(t time.Time) time.Time {
	return GetFirstDayOfMonth(t).AddDate(0, 1, -1)
}

//获取某一天的0点时间
func GetZero(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
