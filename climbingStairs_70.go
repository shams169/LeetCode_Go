package main

/*
You are climbing a staircase. It takes n steps to reach the top.
Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
*/

func climbingStairs_Fibo(n int) int {

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return climbingStairs_Fibo(n-1) + climbingStairs_Fibo(n-2)
}

func climbStairs(n int) int {
	if n <= 3 {
		return n
	}

	a := 2
	b := 3
	ans := b
	for i := 4; i <= n; i++ {
		ans = a + b
		a = b
		b = ans
	}
	return ans

}
