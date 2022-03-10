package validate

import (
	"regexp"
)

func Num(str string) bool {
	return regexp.MustCompile(`^\d+$`).MatchString(str)
}

func Mobile(mobileNum string) bool {
	regular := "^1[3-9]\\d{9}$"
	return regexp.MustCompile(regular).MatchString(mobileNum)
}

func Mail(mail string) bool {
	regular := "^[a-zA-Z0-9_.-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$"
	return regexp.MustCompile(regular).MatchString(mail)
}
