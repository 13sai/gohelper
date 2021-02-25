package gohelper

import (
	"fmt"
	"time"
)

// 2006-01-02 15:04:05 记住这一刻

type AmazeTime struct {
	Time         time.Time
	FirstOfMonth time.Time
	LastOfMonth  time.Time
	FirstOfWeek  time.Time
	LastOfWeek   time.Time
}

var nilTime = time.Time{}

func NewTime(t time.Time) *AmazeTime {
	return &AmazeTime{Time: t}
}

func NewNowTime() *AmazeTime {
	return &AmazeTime{Time: time.Now()}
}

// 获取周一时间
func (t *AmazeTime) GetMonday() time.Time {
	offset := int(time.Monday - t.Time.Weekday())
	if offset > 0 {
		offset = -6
	}

	t.FirstOfWeek = time.Date(t.Time.Year(), t.Time.Month(), t.Time.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	return t.FirstOfWeek
}

//获取传入的时间所在月份的第一天
func (t *AmazeTime) GetFirstDateOfMonth() time.Time {
	if t.FirstOfMonth != nilTime {
		return t.FirstOfMonth
	}
	t.FirstOfMonth = t.Time.AddDate(0, 0, -t.Time.Day()+1)
	return t.FirstOfMonth
}

//获取传入的时间所在月份的最后一天
func (t *AmazeTime) GetLastDateOfMonth() time.Time {
	if t.LastOfMonth != nilTime {
		return t.LastOfMonth
	}
	fmt.Println("2415")
	if t.FirstOfMonth == nilTime {
		t.GetFirstDateOfMonth()
	}
	t.LastOfMonth = t.FirstOfMonth.AddDate(0, 1, -1)
	return t.LastOfMonth
}

//获取某一天的0点时间
func (t *AmazeTime) ToZero() time.Time {
	return time.Date(t.Time.Year(), t.Time.Month(), t.Time.Day(), 0, 0, 0, 0, t.Time.Location())
}
