package main

func main() {

}

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return nil
	}

	set := map[byte]int{}
	for i := 0; i < len(p); i++ {
		set[p[i]]++
	}
	c := len(set)

	res := []int{}
	left := 0
	right := 0
	for right < len(s) {
		for right-left < len(p) {
			if _, ok := set[s[right]]; ok {
				set[s[right]]--
				if set[s[right]] == 0 {
					c--
				}
			}
			right++
		}

		if c == 0 {
			res = append(res, left)
		}

		if _, ok := set[s[left]]; ok {
			if set[s[left]] == 0 {
				c++
			}
			set[s[left]]++
		}
		left++
	}

	return res
}
