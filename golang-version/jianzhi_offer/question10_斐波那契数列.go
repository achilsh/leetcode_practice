package jianzhi_offer

func FibonacciInCycle(n int) int {
	a, b := 0, 1
	if n <= 0 {
		return a
	}
	if n <= 1 {
		return b
	}

	var c int // 0, 1, 1, 2
	for i := 2; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}

// QinWaJump 青蛙跳台阶
func QinWaJump(n int) int {
	if n <= 2 {
		return n
	}

	a, b := 1, 2
	var c int
	for i := 3; i <= n; i++ {
		c = a + b
		a = b
		b = c

	}
	return c
}
