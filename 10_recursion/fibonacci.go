package _0_recursion

func Fibonacci(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
