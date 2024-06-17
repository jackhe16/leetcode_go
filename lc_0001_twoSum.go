/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 */

package main

// @lc code=start

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, x := range nums {
		if p, ok := m[target-x]; ok {
			return []int{p, i}
		}
		m[x] = i
	}
	return nil
}

// @lc code=end
