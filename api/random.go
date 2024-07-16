package secretnote

import (
	"fmt"
	"math/rand"
	"strings"
)

func RandInt(min int64, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}
func RandString(length int) string {
	var sb strings.Builder
	sb.Grow(length)
	for i := 0; i < length; i++ {
		sb.WriteByte(byte(RandInt('a', 'z')))
	}
	return sb.String()
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandString(6))
}
