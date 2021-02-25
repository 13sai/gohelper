package gohelper_test

import (
	"testing"

	"github.com/13sai/gohelper"
)

func TestPost(t *testing.T) {
	http := gohelper.NewHttp("https://hotel.hznmd.com")
	res, err := http.SetMethod("POST").Run()
	t.Log(string(res), err)
}

func TestGet(t *testing.T) {
	http := gohelper.NewHttp("https://www.baidu.com/")
	res, err := http.SetMethod("GET").Run()
	t.Log(string(res), err)
}
