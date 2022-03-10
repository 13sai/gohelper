package str

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"unicode/utf8"

	"github.com/rs/xid"
)

func Len(str string) int {
	return utf8.RuneCountInString(str)
}

func Sub(str string, sub ...int) string {
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

func Combine(str ...string) string {
	var bt bytes.Buffer
	for _, arg := range str {
		bt.WriteString(arg)
	}
	return bt.String()
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// 生成唯一id
func GetUUid() string {
	return xid.New().String()
}

func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func Base64Decode(s string) string {
	b, _ := base64.StdEncoding.DecodeString(s)
	return string(b)
}

func Base64UrlEncode(s string) string {
	return base64.URLEncoding.EncodeToString([]byte(s))
}

func Base64UrlDecode(s string) string {
	b, _ := base64.URLEncoding.DecodeString(s)
	return string(b)
}
