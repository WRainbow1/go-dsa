package main

func maxProduct(nums []int) int {
	curr_min := nums[0]
	curr_max := nums[0]
	max_product := nums[0]

	for _, val := range nums[1:] {
		if val < 0 {
			curr_min, curr_max = curr_max, curr_min
		}

		curr_min = min(val, curr_min * val)
		curr_max = max(val, curr_max * val)

		max_product = max(max_product, curr_max)
	}

	return max_product
}

// func main() {
// 	nums := []int{2,3,-2,4}
// 	maxProduct(nums)
// }