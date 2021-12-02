package http

import (
	"testing"
)

func TestPost(t *testing.T) {
	http := New("https://www.baidu.com")
	res, err := http.Method("POST").Run()
	t.Log(string(res), err)
}

func TestGet(t *testing.T) {
	http := New("https://www.baidu.com/")
	res, err := http.Method("GET").Run()
	t.Log(string(res), err)
}
