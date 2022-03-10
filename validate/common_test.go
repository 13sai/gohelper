package validate

import (
	"testing"
)

func TestNum(t *testing.T) {
	t.Logf("num233-%v-%v-%v", Num("233"), "23s3", Num("23s3"))
}

func TestMobile(t *testing.T) {
	t.Logf("mobile12345678901-%v-%v-%v", Mobile("12345678901"), "13333333333", Mobile("13333333333"))
}

func TestMail(t *testing.T) {
	t.Logf("mail1234@22-%v-%v-%v", Mail("1234@22"), "4321@qq.com", Mail("4321@qq.com"))
}
