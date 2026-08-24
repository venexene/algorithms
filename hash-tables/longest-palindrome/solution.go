package main

func main() {

}

func longestPalindrome(s string) int {
	letters := make(map[rune]struct{}, len(s))

	res := 0
	for _, r := range s {
		if _, ok := letters[r]; ok {
			delete(letters, r)
			res += 2
		} else {
			letters[r] = struct{}{}
		}
	}
	if len(letters) == 0 {
		return res
	}
	return res + 1
}
