package Array

/*
leetcode: https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/
*/

/*
solution1:
固定left, 在剩余区间进行二分查找
时间复杂度 O(nLogN) 空间 O(1)
*/
func twoSum1(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for index, item := range numbers {
		remain := target - item
		left = index + 1

		// 二分查找
		res, ok := func(numbers []int, target, left, right int) (int, bool) {
			for left <= right {
				mid := left + (right-left)>>1
				if numbers[mid] > target {
					right = mid - 1
				} else if numbers[mid] < target {
					left = mid + 1
				} else {
					return mid, true
				}
			}
			return 0, false
		}(numbers, remain, left, right)

		if ok {
			//题目要求 下标从1开始
			return []int{index + 1, res + 1}
		}
	}

	return nil
}

/*
solution2:
双指针
时间O(N) 空间O(1)
*/
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}
