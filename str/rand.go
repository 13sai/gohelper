package str

import (
	"math/rand"
	"time"
	"unsafe"
)

const (
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

var randS = rand.NewSource(time.Now().UnixNano())

func Random(n int, s string) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, randS.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = randS.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(s) {
			b[i] = s[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}
