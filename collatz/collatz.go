package collatz

func Collatz(n uint64) []uint64 {
	var resutls []uint64
	r := n
	for {
		resutls = append(resutls, r)
		if r == 1 {
			break
		}
		r = collatz(r)
	}

	return resutls
}

func Collatzs(start, end uint64) [][]uint64 {
	var results [][]uint64
	for i := start; i <= end; i++ {
		results = append(results, Collatz(i))
	}

	return results
}

func collatz(n uint64) uint64 {
	if (n & 1) == 1 {
		return n*3 + 1
	} else {
		return n / 2
	}
}
