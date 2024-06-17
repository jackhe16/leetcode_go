/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 */

package main

// @lc code=start

func removeElement(nums []int, val int) int {
	k := 0
	fast := 0

	for ; fast < len(nums); fast = fast + 1 {
		if nums[fast] != val {
			nums[k] = nums[fast]
			k = k + 1
		}

	}

	return k
}

// @lc code=end
