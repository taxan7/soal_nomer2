package main

func Bilanganprima(n int) []int {
	var result []int

	for b1 := 1; b1 <= n; b1++ {
		x := 0
		for b2 := 1; b2 <= b1; b2++ {
			if b1%b2 == 0 {
				x++
			}
		}
		if x == 2 {
			result = append(result, b1)
		}
	}

	return result
}
