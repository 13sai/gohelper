package gohelper_test

import (
	"testing"

	"github.com/13sai/gohelper"
)

func TestStrLen(t *testing.T) {
	t.Log("hello!", gohelper.StrLen("hello!"))
	t.Log("你好啊！", gohelper.StrLen("你好啊！"))
}

func TestStrSub(t *testing.T) {
	t.Log("hello", gohelper.StrSub("hello", 1))
	t.Log("你好啊！", gohelper.StrSub("你好啊！", 2, 10))
}

func TestStrCombine(t *testing.T) {
	t.Log("hello+world=", gohelper.StrCombine("hello", "world"))
	t.Log("你好啊+13sai=", gohelper.StrCombine("你好啊", "13sai"))
}

func TestMd5(t *testing.T) {
	t.Log("123456", gohelper.Md5("123456"))
}

func TestGetUUid(t *testing.T) {
	t.Log("getUUid", gohelper.GetUUid())
}
