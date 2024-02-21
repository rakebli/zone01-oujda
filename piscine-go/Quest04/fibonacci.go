package piscine

func Fibonacci(index int) int {
	var i int

	i = 0
	if index < 0 {
		return (-1)
	} else if index < 2 {
		return (index)
	}
	i = Fibonacci(index-1) + Fibonacci(index-2)
	return i
}
