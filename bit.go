func newBIT(n int) []int {
	return make([]int, n+1)
}
func lowbit(x int) int { return x & -x }
func add(bit []int, i, x, n int) {
	for i <= n {
		bit[i] += x
		i += lowbit(i)
	}
}
func sum(bit []int, i int) int {
	var ret int = 0
	for i > 0 {
		ret += bit[i]
		i -= lowbit(i)
	}
	return ret
}
func rangeSum(bit []int, l, r int) int {
	return sum(bit, r) - sum(bit, l-1)
}
