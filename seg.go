type Node struct {
	x int
}
type segment struct {
	n, m int
	data []Node
	e    Node
	op   func(a, b Node) Node
}

func newsegtree(arr []Node, e Node, op func(a, b Node) Node) *segment {
	var n int = 1
	for n < len(arr) {
		n *= 2
	}
	data := make([]Node, 2*n)
	for i := 0; i < 2*n; i++ {
		data[i] = e
	}
	for i := 0; i < len(arr); i++ {
		data[n+i] = arr[i]
	}
	for i := n - 1; i >= 0; i-- {
		data[i] = op(data[2*i], data[2*i+1])
	}
	return &segment{n, len(arr), data, e, op}
}
func (seg *segment) update(pos int, val Node) {
	pos += seg.n
	seg.data[pos] = val
	for pos > 1 {
		pos /= 2
		seg.data[pos] = seg.op(seg.data[pos*2], seg.data[pos*2+1])
	}
}
func (seg *segment) query(l, r int) Node {
	l, r = l+seg.n, r+seg.n
	var resl, resr Node = seg.e, seg.e
	for l < r {
		if l&1 == 1 {
			resl = seg.op(resl, seg.data[l])
			l++
		}
		if r&1 == 1 {
			r--
			resr = seg.op(seg.data[r], resr)
		}
		l, r = l/2, r/2
	}
	return seg.op(resl, resr)
}
func (seg *segment) maxRight(l int, f func(Node) bool) int {
	if l == seg.m {
		return l
	}
	l += seg.n
	var sm Node = seg.e
	for {
		for l&1 == 0 {
			l /= 2
		}
		if !f(seg.op(sm, seg.data[l])) {
			for l < seg.n {
				l *= 2
				if f(seg.op(sm, seg.data[l])) {
					sm = seg.op(sm, seg.data[l])
					l++
				}
			}
			return l - seg.n
		}
		sm = seg.op(sm, seg.data[l])
		l++
		if l&-l == l {
			return seg.m
		}
	}
}
func (seg *segment) minLeft(r int, f func(Node) bool) int {
	if r == 0 {
		return r
	}
	r += seg.n
	var sm Node = seg.e
	for {
		r--
		for r > 1 && r&1 == 1 {
			r /= 2
		}
		if !f(seg.op(seg.data[r], sm)) {
			for r < seg.n {
				r = 2*r + 1
				if f(seg.op(seg.data[r], sm)) {
					sm = seg.op(seg.data[r], sm)
					r--
				}
			}
			return r + 1 - seg.n
		}
		sm = seg.op(seg.data[r], sm)
		if r&-r == r {
			return 0
		}
	}
}
