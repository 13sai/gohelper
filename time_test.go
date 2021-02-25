package gohelper_test

import (
	"testing"

	"github.com/13sai/gohelper"
)

var t *gohelper.AmazeTime

func init() {
	t = gohelper.NewNowTime()
}

func TestGetMonday(test *testing.T) {
	test.Log(t.GetMonday().Format("2006/01/02"))
}

func TestGetFirstDateOfMonth(test *testing.T) {
	test.Log(t.GetFirstDateOfMonth())
}

func TestGetLastDateOfMonth(test *testing.T) {
	test.Log(t.GetLastDateOfMonth())
}

func TestToZero(test *testing.T) {
	test.Log(t.ToZero())
}
