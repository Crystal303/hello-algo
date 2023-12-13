// File: max_product_cutting_test.go
// Created Time: 2023-07-23
// Author: Reanon (793584285@qq.com)

package chapter_greedy

import (
	"fmt"
	"testing"
)

func TestMaxProductCutting(t *testing.T) {
	n := 58
	// 贪心算法
	res := maxProductCutting(n)
	fmt.Println("最大切分乘积为", res)

	res = maxProductCuttingV2(n)
	fmt.Println("最大切分乘积为", res)
}

func maxProductCuttingV2(n int) int {
	if n <= 3 {
		return n * (n - 1)
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	maxF := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	for i := 4; i <= n; i++ {
		m := 0
		for j := 1; j < i; j++ {
			m = maxF(m, dp[i-j]*j)
		}
		dp[i] = m
	}
	return dp[n]
}
