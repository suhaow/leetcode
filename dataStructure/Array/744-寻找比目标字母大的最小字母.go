package Array

/*
https://leetcode-cn.com/problems/find-smallest-letter-greater-than-target/
*/

/*
思路：
考察寻找左侧边界得二分搜索
*/

func nextGreatestLetter(letters []byte, target byte) byte {
	left, right := 0, len(letters)-1

	for left <= right {
		mid := left + (right-left)>>1
		if letters[mid] <= target {
			left = mid + 1
		} else if letters[mid] > target {
			right = mid - 1
		}
	}

	if left > len(letters)-1 {
		left = 0
	}
	return letters[left]
}
