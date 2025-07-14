package main

import (
	// "fmt"
	"sync"
	// "time"
)

type result struct {
	palin  string
	length int
}

func longestPalindrome(s string) string {

	var wg sync.WaitGroup

	results := make(chan result, 2*len(s))

	for i := range s {

		wg.Add(1)
		go search_palindromes_odd(s, i, &wg, results)

		if i > 0 {
			wg.Add(1)
			go search_palindromes_even(s, i, &wg, results)
		}
	}

	wg.Wait()
	close(results)

	max_len := 0
	var final_palin string
	for out_result := range results {
		if out_result.length > max_len {
			final_palin = out_result.palin
			max_len = out_result.length
		}
	}

	return final_palin
}

func search_palindromes_odd(s string, i int, wg *sync.WaitGroup, results chan<- result) {

	defer wg.Done()

	back := i - 1
	forward := i + 1

	palin := string(s[i])
	length := 1

	for back >= 0 && forward <= len(s)-1 {

		back_r := s[back]
		forward_r := s[forward]

		if back_r == forward_r {
			palin = string(back_r) + palin + string(forward_r)
			length += 2
			back--
			forward++
		} else {
			break
		}
	}

	results <- result{
		palin:  palin,
		length: length,
	}
}

func search_palindromes_even(s string, i int, wg *sync.WaitGroup, results chan<- result) {

	defer wg.Done()
	runes := []rune(s)

	back := i - 2
	forward := i + 1

	if runes[i-1] != runes[i] {
		return
	}

	palin := string(runes[i-1]) + string(runes[i])
	length := 2

	for back >= 0 && forward <= len(s)-1 {

		back_r := s[back]
		forward_r := s[forward]

		if back_r == forward_r {
			palin = string(back_r) + palin + string(forward_r)
			length += 2
			back--
			forward++
		} else {
			break
		}
	}

	results <- result{
		palin:  palin,
		length: length,
	}
}

// func main() {
// 	input := "babad"
// 	start := time.Now()
// 	palin := longestPalindrome(input)
// 	elapsed := time.Since(start)
// 	fmt.Println(palin)
// 	fmt.Println("Time taken:", elapsed)
// }
