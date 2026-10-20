package main

//weightUF

type uf struct {
	n    int
	root []int
	diff []int64
}

func newuf(n int) *uf {
	root := make([]int, n)
	for i := 0; i < n; i++ {
		root[i] = -1
	}
	return &uf{n, root, make([]int64, n)}
}
func (u *uf) find(x int) int {
	p := u.root[x]
	if p < 0 {
		return x
	}
	r := u.find(p)
	u.diff[x] += u.diff[p]
	u.root[x] = r
	return r
}
func (u *uf) weight(x int) int64 {
	u.find(x)
	return u.diff[x]
}
func (u *uf) union(x, y int, w int64) bool {
	w += u.weight(x) - u.weight(y)
	x, y = u.find(x), u.find(y)
	if x == y {
		return w == 0
	}
	if -u.root[x] < -u.root[y] {
		x, y = y, x
		w = -w
	}
	u.root[x] += u.root[y]
	u.root[y] = x
	u.diff[y] = w
	return true
}
func (u *uf) diffOf(x, y int) int64 {
	return u.weight(y) - u.weight(x)
}
func (u *uf) size(x int) int {
	return -u.root[u.find(x)]
}
