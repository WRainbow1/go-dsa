package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
    
	solved := make([]bool, len(s))

	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	for i := range s {
		for j := 0; j <= i; j++ {

			potential_word := s[j:i+1]

			if (j == 0 || solved[j-1]) && wordSet[potential_word] {
				solved[i] = true
				break
			}
		}
	}

	return solved[len(solved)-1]
}

// func main() {
// 	inputS := "a"
// 	inputArr := []string{"a"}
// 	output := wordBreak(inputS, inputArr)
// 	fmt.Println(output)
// }