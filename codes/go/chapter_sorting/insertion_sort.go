// File: insertion_sort.go
// Created Time: 2022-12-12
// Author: msk397 (machangxinq@gmail.com)

package chapter_sorting

/* 插入排序 */
func insertionSort(nums []int) {
	// 外循环：未排序区间为 [0, i]
	for i := 1; i < len(nums); i++ {
		base := nums[i]
		j := i - 1
		// 内循环：将 base 插入到已排序部分的正确位置
		for j >= 0 && nums[j] > base {
			nums[j+1] = nums[j] // 将 nums[j] 向右移动一位
			j--
		}
		nums[j+1] = base // 将 base 赋值到正确位置
	}
}

/* 插入排序 从后往前*/
func insertionSortV2(nums []int) {
	if len(nums) < 1 {
		return
	}
	for i := len(nums) - 2; 0 <= i; i-- {
		base := nums[i]
		j := i
		// 往前挪
		for j < len(nums)-1 && nums[j+1] < base {
			nums[j] = nums[j+1]
			j++
		}
		nums[j] = base
	}
}
