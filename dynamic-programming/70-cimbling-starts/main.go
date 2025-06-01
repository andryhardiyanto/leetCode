package main

import "fmt"

func memorization(n int, memo map[int]int) int {
	if n == 1 || n == 0 {
		return 1
	}
	if val, exists := memo[n]; exists {
		return val
	}

	memo[n] = memorization(n-1, memo) + memorization(n-2, memo)
	return memo[n]
}

func climbStairs(n int) int {
	memo := make(map[int]int, 0)
	return memorization(n, memo)
}

func main() {
	fmt.Println(climbStairs(3))
}
