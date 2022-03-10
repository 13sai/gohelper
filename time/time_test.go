package time

import (
	"testing"
	"time"
)

func TestDateTimeToTimeStamp(t *testing.T) {
	now := "2022-01-01 10"
	t.Logf("DateTimeToTimeStamp=%v", DateTimeToTimeStamp(now, "2006-01-02 15"))
}

func TestTimeStampToDateTime(t *testing.T) {
	t.Logf("TimeStampToDateTime=%v", TimeStampToDateTime(1646890797, "2006-01-02"))
}

func TestGetMonday(t *testing.T) {
	now := time.Now()
	t.Logf("GetMonday=%v", GetMonday(now))
}

func TestGetSunday(t *testing.T) {
	now := time.Now()
	t.Logf("GetMonday=%v", GetMonday(now))
}

func TestGetFirstDayOfMonth(t *testing.T) {
	now := time.Now()
	t.Logf("GetFirstDayOfMonth=%v", GetFirstDayOfMonth(now))
}

func TestGetLastDayOfMonth(t *testing.T) {
	now := time.Now()
	t.Logf("GetLastDayOfMonth=%v", GetLastDayOfMonth(now))
}

func TestGetFirstDayOfYear(t *testing.T) {
	now := time.Now()
	t.Logf("GetFirstDayOfYear=%v", GetFirstDayOfYear(now))
}

func TestGetLastDayOfYear(t *testing.T) {
	now := time.Now().AddDate(0, -4, 0)
	t.Logf("GetLastDayOfYear=%v", GetLastDayOfYear(now))
}

func TestGetZero(t *testing.T) {
	now := time.Now()
	t.Logf("GetZero=%v", GetZero(now))
}
