package main

import (
	"fmt"
)

func main() {
	s := "123"
	res := myAtoi(s)
	fmt.Println(res)
}

func myAtoi(s string) int {
	res := 0
	i := 0
	sign := 1
	for i < len(s) && s[i] == ' ' {
		i++
	}

	if i >= len(s) {
		return 0
	}

	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		i++
	}

	minVal := -1 << 31
	maxVal := 1 << 31 - 1
	for i < len(s) {
		if s[i] < '0' || s[i] > '9' {
			return res * sign
		}

		if res > maxVal/10 {
			if sign == 1 {
				return maxVal
			} else {
				return minVal
			}
		}
		res = res*10 + int(s[i]-'0')
		i++
	}
	res *= sign

	if res > maxVal {
		return maxVal
	} else if res < minVal {
		return minVal
	}
	return res
}
