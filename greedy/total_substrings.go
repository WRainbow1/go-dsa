package main

import (
	"sync"
	"sync/atomic"
)

func countSubstrings(s string) int {

	var wg sync.WaitGroup
	var total_palins int32

	for i := range s {

		wg.Add(1)
		go count_palindromes_odd(s, i, &wg, &total_palins)

		if i > 0 {
			wg.Add(1)
			go count_palindromes_even(s, i, &wg, &total_palins)
		}
	}

	wg.Wait()

	out := int(total_palins)

	return out
}

func count_palindromes_odd(s string, i int, wg *sync.WaitGroup, total_palins *int32) {

	defer wg.Done()

	back := i - 1
	forward := i + 1

	atomic.AddInt32(total_palins, 1)

	for back >= 0 && forward <= len(s)-1 {

		back_r := s[back]
		forward_r := s[forward]

		if back_r == forward_r {
			back--
			forward++
			atomic.AddInt32(total_palins, 1)
		} else {
			break
		}
	}
}

func count_palindromes_even(s string, i int, wg *sync.WaitGroup, total_palins *int32) {

	defer wg.Done()
	runes := []rune(s)

	back := i - 2
	forward := i + 1

	if runes[i-1] != runes[i] {
		return
	}

	atomic.AddInt32(total_palins, 1)

	for back >= 0 && forward <= len(s)-1 {

		back_r := s[back]
		forward_r := s[forward]

		if back_r == forward_r {
			back--
			forward++
			atomic.AddInt32(total_palins, 1)
		} else {
			break
		}
	}
}

// func main() {
// 	input := "aaa"
// 	start := time.Now()
// 	total_palins := countSubstrings(input)
// 	elapsed := time.Since(start)
// 	fmt.Println(total_palins)
// 	fmt.Println("Time taken:", elapsed)
// }
