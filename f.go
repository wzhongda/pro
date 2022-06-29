package main

import "fmt"

func a() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	taget := 3
	len := len(nums)
	left := 0
	right := len - 1
	middle := 0
	for {
		if left <= right {
			middle = left + ((right - left) / 2)
			if nums[middle] > taget {
				right = middle - 1
			} else if nums[middle] < taget {
				left = middle + 1
			} else {
				break
			}
		}
	}
	fmt.Println(nums[middle])

}

func main() {
	a()
}
