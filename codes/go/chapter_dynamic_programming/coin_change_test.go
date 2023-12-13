// File: coin_change_test.go
// Created Time: 2023-07-23
// Author: Reanon (793584285@qq.com)

package chapter_dynamic_programming

import (
	"fmt"
	"math"
	"testing"
)

func TestCoinChange(t *testing.T) {
	coins := []int{1, 2, 5}
	amt := 6

	// 动态规划
	res := coinChangeDP(coins, amt)
	fmt.Printf("凑到目标金额所需的最少硬币数量为 %d\n", res)

	// 空间优化后的动态规划
	res = coinChangeDPComp(coins, amt)
	fmt.Printf("凑到目标金额所需的最少硬币数量为 %d\n", res)

	res = coinChange(coins, amt)
	fmt.Printf("凑到目标金额所需的最少硬币数量为 %d\n", res)
}

func coinChange(coins []int, amt int) int {
	res := make([][]int, 0)
	state := make([]int, 0)
	selected := make([]bool, len(coins))
	coinChangeBack(coins, amt, &state, &res, selected)

	times := math.MaxInt
	for i := range res {
		if 0 < len(res[i]) && len(res[i]) < times {
			times = len(res[i])
		}
	}
	return times
}

func coinChangeBack(coins []int, amt int, state *[]int, res *[][]int, selected []bool) {
	if amt == 0 {
		tmp := make([]int, len(*state))
		copy(tmp, *state)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < len(coins); i++ {
		if amt-coins[i] < 0 {
			continue
		}
		if !selected[i] {
			selected[i] = true
			*state = append(*state, coins[i])
			coinChangeBack(coins, amt-coins[i], state, res, selected)
			*state = (*state)[:len(*state)-1]
			selected[i] = false
		}
	}
}
