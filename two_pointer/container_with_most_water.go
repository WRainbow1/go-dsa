package main

func maxArea(height []int) int {

	max_area := 0
	area := 0
	left_pointer := 0
	right_pointer := len(height) - 1

	for {
		left_height := height[left_pointer]
		right_height := height[right_pointer]

		if left_pointer == right_pointer {
			break
		}

		if left_height >= right_height {
			area = right_height * (right_pointer - left_pointer)
			max_area = checkMax(area, max_area)
			right_pointer--
		} else if right_height > left_height {
			area = left_height * (right_pointer - left_pointer)
			max_area = checkMax(area, max_area)
			left_pointer++
		}
	}
	return max_area
}

func checkMax(curr int, max int) int {
	if curr > max {
		return curr
	} else {
		return max
	}
}

// func main() {
// 	heights := []int{1,8,6,2,5,4,8,3,7,2,4,3,4,20}
// 	max := maxArea(heights)
// 	fmt.Println(max)
// }
