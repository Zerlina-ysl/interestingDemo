package sorting_algorithm

import (
	"fmt"
	"testing"
)

func Test_Quick(t *testing.T) {
	test := []int{1, 9, 101, 88, 33, 2, 56, 67, 432, 23, 129, 98, 1298}
	QuickSort(test)
	fmt.Println(test)
}
func QuickSort(nums []int) {
	var quick func(nums []int, left, right int)
	quick = func(nums []int, left, right int) {
		if left > right {
			return
		}
		i, j, pivot := left, right, nums[left]
		for i < j {
			for i < j && pivot <= nums[j] {
				j--
			}
			for i < j && pivot >= nums[i] {
				i++
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		nums[i], nums[left] = nums[left], nums[i]
		quick(nums, left, i-1)
		quick(nums, i+1, right)
		return

	}
	quick(nums, 0, len(nums)-1)
}
