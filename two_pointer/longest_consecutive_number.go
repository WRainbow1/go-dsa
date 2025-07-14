package main

// import "fmt"

func longestConsecutive(nums []int) int {
	
	if i := len(nums); i == 0 {
		return 0
	}

	max_consec := 1

    set := make(map[int]struct{})
	for _, val := range nums {
		set[val] = struct{}{}
	}

    for num := range set {

        if _, exists := set[num-1]; !exists {

            current_val := num
            current_consec := 1
            for {
                target_val := current_val + 1
                if _, exists := set[target_val]; exists {
                    current_consec += 1
                    current_val += 1
                    if current_consec > max_consec {
                        max_consec = current_consec
                    }
                } else {
                    break
                }
            }
        }
    }
	return max_consec
}

// func main() {
//     nums := []int{1,0,1,2}
//     result := longestConsecutive(nums)
//     fmt.Println(result)
// }