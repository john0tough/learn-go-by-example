package iteration

import "strings"

const repeatCount = 5

func Repeat(char string) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(char)
	}
	return repeated.String()
}
