// File: edit_distance_test.go
// Created Time: 2023-07-23
// Author: Reanon (793584285@qq.com)

package chapter_dynamic_programming

import (
	"fmt"
	"testing"
)

func TestEditDistanceDFS(test *testing.T) {
	s := "bag"
	t := "jpack"
	n := len(s)
	m := len(t)

	// 暴力搜索
	res := editDistanceDFS(s, t, n, m)
	fmt.Printf("将 %s 更改为 %s 最少需要编辑 %d 步\n", s, t, res)

	// 记忆化搜索
	mem := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		mem[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			mem[i][j] = -1
		}
	}
	res = editDistanceDFSMem(s, t, mem, n, m)
	fmt.Printf("将 %s 更改为 %s 最少需要编辑 %d 步\n", s, t, res)

	// 动态规划
	res = editDistanceDP(s, t)
	fmt.Printf("将 %s 更改为 %s 最少需要编辑 %d 步\n", s, t, res)

	// 空间优化后的动态规划
	res = editDistanceDPComp(s, t)
	fmt.Printf("将 %s 更改为 %s 最少需要编辑 %d 步\n", s, t, res)

	res = editDistanceDPV1(s, t)
	fmt.Printf("将 %s 更改为 %s 最少需要编辑 %d 步\n", s, t, res)
	res = editDistanceDFSV1(s, t)
	fmt.Printf("将 %s 更改为 %s 最少需要编辑 %d 步\n", s, t, res)
	res = editDistanceDPCompV1(s, t)
	fmt.Printf("将 %s 更改为 %s 最少需要编辑 %d 步\n", s, t, res)
}

func editDistanceDFSV1(s string, t string) int {
	if len(t) == 0 {
		return len(s)
	}
	if len(s) == 0 {
		return len(t)
	}
	if s[len(s)-1] == t[len(t)-1] {
		return editDistanceDFSV1(s[:len(s)-1], t[:len(t)-1])
	}
	minFunc := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	add := editDistanceDFSV1(s[:len(s)-1], t)
	del := editDistanceDFSV1(s, t[:len(t)-1])
	repl := editDistanceDFSV1(s[:len(s)-1], t[:len(t)-1])
	return minFunc(minFunc(add, del), repl) + 1
}

func editDistanceDPV1(s string, t string) int {
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
		dp[i][0] = i
	}
	for j := range dp[0] {
		dp[0][j] = j
	}

	minFunc := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1]
				continue
			}
			// dp[i][j-1] 添加
			// dp[i-1][j] 删除
			// dp[i-1][j-1] 替换
			dp[i][j] = minFunc(minFunc(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1]) + 1
		}
	}

	return dp[len(s)][len(t)]
}

func editDistanceDPCompV1(s string, t string) int {
	dp := make([]int, len(t)+1)
	for j := range dp {
		dp[j] = j
	}

	minFunc := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}

	for i := 1; i <= len(s); i++ {
		leftup := dp[0]
		dp[0] = i

		for j := 1; j <= len(t); j++ {
			tmp := dp[j]
			if s[i-1] == t[j-1] {
				dp[j] = leftup
			} else {
				dp[j] = minFunc(minFunc(leftup, dp[j-1]), dp[j]) + 1
			}
			leftup = tmp
		}
	}

	return dp[len(t)]
}
