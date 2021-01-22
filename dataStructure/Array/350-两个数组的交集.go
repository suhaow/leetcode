package Array

import (
	"sort"
)

/*
leetcode: https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/
*/

/*
笨比操作：没审题 以为是排序的。。
先排序
对于长度较小的数组进行遍历，挨个判断当前元素是否在另一个数组中存在，若存在删除两个数组中该元素
*/

func intersect1(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	if len(nums2) < len(nums1) {
		tmp := nums1
		nums1 = nums2
		nums2 = tmp
	}
	var res []int
	for _, item := range nums1 {
		var ok bool
		nums2, ok = binaryRemove(item, nums2)
		if ok {
			res = append(res, item)
		}
	}
	return res
}

/*
删除第一个与key相等的元素
*/
func binaryRemove(key int, nums []int) ([]int, bool) {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)>>1
		//相等 则删除nums[mid]
		if nums[mid] == key {
			var tmp []int
			tmp = append(tmp, nums[0:mid]...)
			tmp = append(tmp, nums[mid+1:]...)
			return tmp, true
		} else if nums[mid] > key {
			right = mid - 1
		} else if nums[mid] < key {
			left = mid + 1
		}
	}
	return nums, false
}

/*
哈希表
*/
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums2) < len(nums1) {
		tmp := nums1
		nums1 = nums2
		nums2 = tmp
	}

	//将
	numMap := make(map[int]int)
	for _, item := range nums1 {
		numMap[item]++
	}

	var res []int
	for _, item := range nums2 {
		if numMap[item] > 0 {
			res = append(res, item)
			numMap[item]--
		}
	}

	return res
}
