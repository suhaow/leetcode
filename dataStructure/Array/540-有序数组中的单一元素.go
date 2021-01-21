package Array

/*
leetcode:https://leetcode-cn.com/problems/single-element-in-a-sorted-array/
note: https://leetcode-cn.com/problems/single-element-in-a-sorted-array/solution/you-xu-shu-zu-zhong-de-dan-yi-yuan-su-by-leetcode/
*/

/*
给定一个只包含整数的有序数组，每个元素都会出现两次，唯有一个数只会出现一次，找出这个数。

示例 1:

输入: [1,1,2,3,3,4,4,8,8]
输出: 2
示例 2:

输入: [3,3,7,7,10,11,11]
输出: 10

注意: 您的方案应该在 O(log n)时间复杂度和 O(1)空间复杂度中运行。
*/

/*
二分查找法
中间元素的同一元素在右侧，且被mid分成两半的数组为偶数(mid下标为偶数) left=mid+2    例如 1 1 4 4 5(mid) 5 6 8 8
中间元素的同一元素在右边，且报mid分成两半的数组为基数(mid下标为偶数) right=mid-1   例如 1 2 2 3(mid) 3 4 4

中间元素的同一元素在左侧，且被mid分成两半的数组为偶数(mid下标为偶数) right=mid-2    例如 1 1 4 5 5(mid) 6 6 8 8
中间元素的同一元素在左侧，且报mid分成两半的数组为基数(mid下标为偶数) left=mid-1    例如 1 1 2 2(mid) 3 4 4
*/
func singleNonDuplicateBinarySearch(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		index := mid % 2
		if nums[mid] == nums[mid+1] {
			if index == 0 {
				left = mid + 2
			} else {
				right = mid - 1
			}
		} else if nums[mid] == nums[mid-1] {
			if index == 0 {
				right = mid - 2
			} else {
				left = mid + 1
			}
		} else {
			return nums[mid]
		}
	}
	return nums[left]
}

/*
暴力法：顺序遍历
判断第一个元素和第二个元素是否相等，不同则说明第一个元素是单一元素
如果达到最后一个元素，则视为单一元素
*/
func singleNonDuplicateViolence(nums []int) int {
	for i := 0; i < len(nums)-1; i += 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}

/*
仅对偶数索引进行二分搜索
*/
func singleNonDuplicate(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)>>1
		if mid%2 == 1 {
			mid--
		}

		if nums[mid] == nums[mid+1] {
			left = mid + 2
		} else {
			right = mid
		}
	}

	//单个元素直接返回
	return nums[left]
}
