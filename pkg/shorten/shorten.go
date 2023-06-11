package shorten

import (
	"math/rand"
	"strings"
	"time"
)

var chars string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"

func Shorten(url string) string {
	var builder strings.Builder
	builder.Grow(8)
	rand.Seed(time.Now().UnixNano())
	var alphabet []rune = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890")

	builder.WriteString(encode(time.Now().Unix(), 3))
	builder.WriteString(randomString(5, alphabet))

	return builder.String()
}

func encode(num int64, bit int) string {
	bytes := []byte{}
	for i := 0; i < bit; i++ {
		bytes = append(bytes, chars[num%62])
		num = num / 62
	}
	// return bytes
	reverse(bytes)
	return string(bytes)
}

func reverse(a []byte) {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
}

func randomString(n int, alphabet []rune) string {

	alphabetSize := len(alphabet)
	var sb strings.Builder

	for i := 0; i < n; i++ {
		ch := alphabet[rand.Intn(alphabetSize)]
		sb.WriteRune(ch)
	}

	s := sb.String()
	return s
}
