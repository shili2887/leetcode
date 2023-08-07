package main

import (
	"fmt"
	"math"
)

func findMedianSortedArray(nums []int) float64 {
	n := len(nums)
	if n == 0 {
		return 0.0
	}
	if math.Mod(float64(n), 2) == 1 {
		return float64(nums[n/2])
	}

	return (float64(nums[n/2-1]) + float64(nums[n/2])) / 2.0

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m == 0 {
		return findMedianSortedArray(nums2)
	} else if n == 0 {
		return findMedianSortedArray(nums1)
	}
	if m > n {
		tempArray := nums1
		nums1 = nums2
		nums2 = tempArray

		temp := n
		n = m
		m = temp
	}
	min := 0
	max := m
	halfLen := (m + n + 1) / 2

	for min <= max {
		i := (min + max) / 2
		j := halfLen - i
		if i < max && nums2[j-1] > nums1[i] {
			min++
		} else if i > min && nums1[i-1] > nums2[j] {
			max--
		} else {
			var maxLeft, minRight float64

			if i == 0 {
				maxLeft = float64(nums2[j-1])
			} else if j == 0 {
				maxLeft = float64(nums1[i-1])
			} else {
				maxLeft = math.Max(float64(nums1[i-1]), float64(nums2[j-1]))
			}

			if i == m {
				minRight = float64(nums2[j])
			} else if j == n {
				minRight = float64(nums1[i])
			} else {
				minRight = math.Min(float64(nums1[i]), float64(nums2[j]))
			}

			if math.Mod(float64(m+n), 2) == 1 {
				return maxLeft
			}
			return (minRight + maxLeft) / 2.0
		}
	}
	return 0.0
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
