package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := make([]int, len(nums1)+len(nums2))
	copy(merged, nums1)
	copy(merged[len(nums1):], nums2)

	sort.Ints(merged)
	if len(merged)%2 == 0 {
		mid1 := len(merged) / 2
		mid2 := mid1 - 1
		return float64(merged[mid1]+merged[mid2]) / 2.0
	}
	mid := len(merged) / 2
	return float64(merged[mid])
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2, 3}
	median := findMedianSortedArrays(nums1, nums2)
	fmt.Println(median)
}
