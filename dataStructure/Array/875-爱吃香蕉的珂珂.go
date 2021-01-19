package Array

/*
note: https://mp.weixin.qq.com/s/QC24hyg0ZgjR7-LgnEzMYg
leetcode: https://leetcode-cn.com/problems/koko-eating-bananas/
*/

/*
暴力解法：速度初始值为1，依次增长。找到第一个满足时间限制的
*/
func minEatingSpeed1(piles []int, H int) int {
	maxSpeed := getMax(piles)
	for speed := 1; speed <= maxSpeed; speed++ {
		if canFinish(piles, speed, H) {
			return speed
		}
	}
	return 0
}

/*
搜索左侧边界的二分查找来替代暴力中的顺序遍历
*/
func minEatingSpeed(piles []int, H int) int {
	right := getMax(piles) + 1
	left := 1
	//不断缩小左边界
	for left < right {
		mid := left + (right-left)>>1
		if canFinish(piles, mid, H) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func canFinish(piles []int, speed, H int) bool {
	time := 0
	for _, item := range piles {
		time += timeOf(item, speed)
	}
	return time <= H
}

func timeOf(n, speed int) int {
	if n%speed > 0 {
		return n/speed + 1
	}
	return n / speed
}

func getMax(piles []int) int {
	max := 0

	for _, item := range piles {
		if item > max {
			max = item
		}
	}

	return max
}
