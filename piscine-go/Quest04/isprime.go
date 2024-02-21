package piscine

func IsPrime(nb int) bool {
	if nb == 2 {
		return true
		//} //else if nb < 0 {
		//return false
	} else if nb%2 == 0 || nb < 2 {
		return false
	}
	for i := 3; i <= nb/2; i += 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
