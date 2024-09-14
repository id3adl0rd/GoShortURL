package tools

import "strings"

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	capacity = 10
)

func UrlGenerator(number uint64) string {
	l := len(alphabet)
	var builder strings.Builder
	builder.Grow(capacity)
	for ; number > 0; number = number / uint64(l) {
		builder.WriteByte(alphabet[number%uint64(l)])
	}

	return builder.String()
}
