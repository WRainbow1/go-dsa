package main

// import "fmt"

func lengthOfLongestSubstring(s string) int {

	var substring []string
	subset := make(map[string]struct{})

	current_len := 0
	max_len := 0

	for _, r := range s {

		char := string(r)

		if _, exists := subset[char]; !exists {
			substring = append(substring, char)
			subset[char] = struct{}{}
			current_len++
		} else {
			substring = append(substring, char)
			current_len++

			for len(substring) > 1 {

				schar := string(substring[0])

				delete(subset, schar)
				substring = substring[1:]
				current_len -= 1

				if char == schar {
					subset[char] = struct{}{}
					break
				}
			}
		}

		if current_len > max_len {
			max_len = current_len
		}

	}
	return max_len
}

// func main() {
// 	bstring := "abcabcbb"
// 	result := lengthOfLongestSubstring(bstring)
// 	fmt.Println(result)
// }
