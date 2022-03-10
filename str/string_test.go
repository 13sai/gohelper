package str

import (
	"testing"
)

func TestStrLen(t *testing.T) {
	t.Log("hello!", Len("hello!"))
	t.Log("你好啊！", Len("你好啊！"))
}

func TestStrSub(t *testing.T) {
	t.Log("hello", Sub("hello", 1))
	t.Log("你好啊！", Sub("你好啊！", 2, 10))
}

func TestStrCombine(t *testing.T) {
	t.Log("hello+world=", Combine("hello", "world"))
	t.Log("你好啊+13sai=", Combine("你好啊", "13sai"))
}

func TestMd5(t *testing.T) {
	t.Log("123456", Md5("123456"))
}

func TestGetUUid(t *testing.T) {
	t.Log("getUUid", GetUUid())
}
