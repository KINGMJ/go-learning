package digital_root

import (
	"strconv"
	"unicode/utf8"
)

func DigitalRoot(n int) int {
	return digitalAdd(n)
}

func digitalAdd(n int) int {
	var sum int
	str := strconv.Itoa(n)
	for _, str := range str {
		buf := make([]byte, 1)
		_ = utf8.EncodeRune(buf, str)
		value, _ := strconv.Atoi(string(buf))
		sum += value
	}

	if len(strconv.Itoa(sum)) == 1 {
		return sum
	} else {
		return digitalAdd(sum)
	}
}
