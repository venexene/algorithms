package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "1010"
	b := "1011"
	res := addBinary(a, b)
	fmt.Println(res)
}

func addBinary(a string, b string) string {
	buf := make([]byte, 0, max(len(a), len(b))+1)

	i := len(a) - 1
	j := len(b) - 1
	prev := 0
	for i >= 0 || j >= 0 {
		ba := 0
		if i >= 0 {
			ba = int(a[i] - '0')
		}

		bb := 0
		if j >= 0 {
			bb = int(b[j] - '0')
		}

		res := ba + bb + prev
		if res > 1 {
			prev = 1
		} else {
			prev = 0
		}

		buf = append(buf, byte(res%2+'0'))
		i--
		j--
	}

	if prev != 0 {
		buf = append(buf, byte(prev+'0'))
	}

	var sb strings.Builder
	for i := len(buf) - 1; i >= 0; i-- {
		sb.WriteByte(buf[i])
	}

	return sb.String()
}
