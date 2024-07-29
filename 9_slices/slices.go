package main

import (
	"fmt"
)

// slices --> dynamic array
// most used construct in go
// + usefull methods
func main() {
	// unintialized slice is nil
	// var nums []int
	// fmt.Println(nums == nil) // go run 9_slices/slices.go
	// fmt.Println(len(nums))

	// var nums = make([]int, 0, 5)
	// capacity --> maximum numbers of a elements can fit
	// fmt.Println(nums) // go run 9_slices/slices.go
	// fmt.Println(len(nums))
	// fmt.Println(nums == nil)
	// fmt.Println(cap(nums)) // for capacity

	// nums := []int{}
	// nums = append(nums, 1)
	// nums = append(nums, 3)
	// nums = append(nums, 5)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))

	// copy function

	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))

	// nums = append(nums, 2)
	// // copy function
	// //   dest   sours
	// copy(nums2, nums)
	// fmt.Println(nums, nums2)

	// slices operator
	// var nums = []int{1, 2, 3}
	// fmt.Println(nums[0:1])
	// fmt.Println(nums[:2])
	// fmt.Println(nums[1:])

	// slice package
	// var nums1 = []int{1, 2}
	// var nums2 = []int{1, 2}
	// fmt.Println(slices.Equal(nums1, nums2))

	var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums)
}
