package time

import (
	"testing"
	"time"
)

func TestGetMonday(t *testing.T) {
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

func TestGetZero(t *testing.T) {
	now := time.Now()
	t.Logf("GetZero=%v", GetZero(now))
}
