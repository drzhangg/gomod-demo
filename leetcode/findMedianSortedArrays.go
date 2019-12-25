package main

import (
	"fmt"
	"sort"
)

//
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		if len(nums2) % 2 == 0 {
			return float64((nums2[len(nums2)/2] + nums2[len(nums2)/2 - 1])) / 2.0
		} else {
			return float64(nums2[len(nums2)/2])
		}
	}
	if len(nums2) == 0 {
		if len(nums1) % 2 == 0 {
			return float64((nums1[len(nums1)/2] + nums1[len(nums1)/2 - 1])) / 2.0
		} else {
			return float64(nums1[len(nums1)/2])
		}
	}
	nums := append(nums1, nums2...)
	// 对数组进行排序
	sort.Ints(nums)
	if len(nums) % 2 == 0 {
		return float64((nums[len(nums)/2] + nums[len(nums)/2 - 1])) / 2.0
	} else {
		return float64(nums[len(nums)/2])
	}
}

func main() {
	fmt.Println(float64(10.0)/2.0)
}
