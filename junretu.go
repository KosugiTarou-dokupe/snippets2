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
	var rev = func(x []E) []E {
		var ret []E = make([]E, len(x))
		for i := 0; i < len(x); i++ {
			ret[len(x)-i-1] = x[i]
		}
		return ret
	}
	x = append(x[:j], rev(x[j:])...)
	return true
}
