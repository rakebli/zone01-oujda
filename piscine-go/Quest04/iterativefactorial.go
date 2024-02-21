package piscine

func IterativeFactorial(nb int) int {
	n := 1
	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	}
	for i := 1; i <= nb; i++ {
		n *= i
		if n < 0 {
			return 0
		}
	}
	return n
}
