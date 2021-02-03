package Array

/*
note:https://mp.weixin.qq.com/s/ioKXTMZufDECBUwRRp3zaA
leetcode: https://leetcode-cn.com/problems/permutation-in-string/
*/

/*
给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
换句话说，第一个字符串的排列之一是第二个字符串的子串。

输入: s1 = "ab" s2 = "eidbaooo"
输出: True
解释: s2 包含 s1 的排列之一 ("ba").
*/

func checkInclusion(s1 string, s2 string) bool {
	need := make(map[byte]int)
	window := make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}

	left, right := 0, 0
	valid := 0

	for right < len(s2) {
		c := s2[right]
		right++

		//窗口信息维护
		if need[c] > 0 {
			window[c]++
			if need[c] == window[c] {
				valid++
			}
		}

		//左侧窗口缩进条件
		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}

			d := s2[left]
			left++

			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return false
}
