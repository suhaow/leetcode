package Array

/*
note: https://mp.weixin.qq.com/s/QC24hyg0ZgjR7-LgnEzMYg
leetcode:https://leetcode-cn.com/problems/capacity-to-ship-packages-within-d-days/
*/

/*
思路： 利用二分搜索优化搜索空间
*/
func shipWithinDays(weights []int, D int) int {
	/*
		货物不可分割，left是max(weights)
		right 是sum(weights)
	*/
	left := func(weights []int) int {
		max := 0

		for _, item := range weights {
			if item > max {
				max = item
			}
		}
		return max
	}(weights)

	right := func(weights []int) int {
		sum := 0
		for _, item := range weights {
			sum += item
		}
		return sum
	}(weights)

	for left < right {
		mid := left + (right-left)>>1
		if canDone(weights, D, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

// 如果载重为 cap，是否能在 D 天内运完货物？
func canDone(weights []int, D, cap int) bool {
	i := 0
	for day := 0; day < D; day++ {
		maxCap := cap
		for maxCap-weights[i] >= 0 {
			maxCap -= weights[i]
			i++
			if i == len(weights) {
				return true
			}
		}
	}
	return false
}
