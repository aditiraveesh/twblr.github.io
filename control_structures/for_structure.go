package control_structures

func collatzChainLength(n int) int {
	length := 0

	for n != 1 {
		if n % 2 == 0 {
			n /= 2
		} else {
			n = 3 * n + 1
		}
		length ++
	}

	return length
}
