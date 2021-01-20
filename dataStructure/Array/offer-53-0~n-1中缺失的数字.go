package Array

/*
https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/
*/

/*
一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。
在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字

示例 1:

输入: [0,1,3]
输出: 2
示例 2:

输入: [0,1,2,3,4,5,6,7,9]
输出: 8
*/

/*
思路
左子数组中下标与元素值相等，右子数组中不相等
[left, right]区间内，使用nums[i] 与 i进行比较
若 nums[mid] == mid， 则说明缺少的数字在右侧 -> left = mid+1
若 nums[mid] != mid， 则说明缺少的数字在左侧 -> right = mid-1
跳出循环时，left 指向右子数组的首位元素，right 指向左子数组的末尾元素
*/
func missingNumber(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
