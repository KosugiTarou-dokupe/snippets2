package main

// junretu.go
func NextPermutation[E any](x []E, cmp func(a, b E) int) bool {
	if len(x) <= 1 {
		return false
	}
	n := len(x)
	var j int = n - 1
	for cmp(x[j-1], x[j]) >= 0 {
		j--
		if j == 0 {
			break
		}
	}
	if j == 0 {
		return false
	}
	var l int = n - 1
	for cmp(x[j-1], x[l]) >= 0 {
		l--
		if l == 0 {
			break
		}
	}
	x[j-1], x[l] = x[l], x[j-1]
	for l, r := j, n-1; l < r; l, r = l+1, r-1 {
		x[l], x[r] = x[r], x[l]
	}
	return true
}
