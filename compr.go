import "slices"

func compress(n int, a []int64) ([]int, []int64) {
	var b []int64 = make([]int64, n)
	copy(b, a)
	slices.Sort(b)
	b = slices.Compact(b)
	var comp []int = make([]int, n)
	for i := 0; i < n; i++ {
		l, r := 0, len(b)
		for r-l > 1 {
			mid := (r + l) / 2
			if b[mid] <= a[i] {
				l = mid
			} else {
				r = mid
			}
		}
		comp[i] = l
	}
	return comp, b
}
