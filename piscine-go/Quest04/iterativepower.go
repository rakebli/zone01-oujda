package piscine

func IterativePower(nb int, power int) int {
	a := 1
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	}
	for i := 0; i < power; i++ {
		a *= nb
	}
	return a
}
