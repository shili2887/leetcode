package main

import "fmt"

// 现在nums2
func merge(nums1 []int, m int, nums2 []int, n int) []int {
	// 边际条件
	if m == 0 {
		return nums2
	}

	if n == 0 {
		return nums1
	}

	nums3 := []int{}

	i := 0
	j := 0
	for k := 0; k < m+n; k++ {
		fmt.Println(k)
		if i >= m {
			nums3 = append(nums3, nums2[j:]...)
			break
		}

		if j >= n {
			nums3 = append(nums3, nums1[i:]...)
			break
		}

		if nums1[i] < nums2[j] {
			nums3 = append(nums3, nums1[i])
			i++
		} else {
			nums3 = append(nums3, nums2[j])
			j++
		}
	}

	copy(nums1, nums3)
	return nums3
}

// 双指针逆序
func mergeReverse(nums1 []int, m int, nums2 []int, n int) []int {
	i := m - 1
	j := n - 1
	k := m + n - 1

	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			k--
			i--
		} else {
			nums1[k] = nums2[j]
			k--
			j--
		}
	}

	return nums1
}

func main() {
	nums1 := []int{4, 0, 0, 0, 0, 0}
	nums2 := []int{1, 2, 3, 5, 6}
	nums := mergeReverse(nums1, 1, nums2, 5)

	fmt.Println(nums)
}
