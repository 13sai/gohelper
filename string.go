package gohelper

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"unicode/utf8"

	"github.com/rs/xid"
)

// 字符串长度utf8
func StrLen(str string) int {
	return utf8.RuneCountInString(str)
}

// 截取字符串
func StrSub(str string, sub ...int) string {
	start := sub[0]
	length := 0
	if len(sub) > 1 {
		length = sub[1]
	}

	if length < 1 {
		return string(([]rune(str))[start:])
	}
	return string(([]rune(str))[start:length])
}

// 合并字符串
func StrCombine(str ...string) string {
	var bt bytes.Buffer
	for _, arg := range str {
		bt.WriteString(arg)
	}
	//获得拼接后的字符串
	return bt.String()
}

// md5加密
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// 生成唯一id
func GetUUid() string {
	return xid.New().String()
}
