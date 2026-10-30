type Node struct {
	x int
}
type segment struct {
	n, m              int //nはセグ木の長さ、mは元のslice長
	data              []Node
	left, right, root []int32
	e                 Node
	op                func(a, b Node) Node
}

func newsegtree(arr []Node, e Node, op func(a, b Node) Node, q int) *segment {
	//qはupdateの回数
	var n int = 1
	logn := 0
	for n < len(arr) {
		logn++
		n *= 2
	}
	data, left, right, root := make([]Node, 2*n, 2*n+q*(1+logn)), make([]int32, 2*n, 2*n+q*(1+logn)), make([]int32, 2*n, 2*n+q*(1+logn)), []int32{1}
	for i := 0; i < 2*n; i++ {
		data[i] = e
	}
	for i := 1; i < n; i++ {
		left[i], right[i] = int32(2*i), int32(2*i+1)
	}
	for i := 0; i < len(arr); i++ {
		data[n+i] = arr[i]
	}
	for i := n - 1; i >= 1; i-- {
		data[i] = op(data[2*i], data[2*i+1])
	}
	return &segment{n, len(arr), data, left, right, root, e, op}
}
func (seg *segment) update(ver, pos int, val Node) int {
	now, l, r := seg.root[ver], 0, seg.n
	var path []int32
	for {
		path = append(path, now)
		if r-l == 1 {
			break
		}
		mid := (l + r) / 2
		if pos < mid {
			now = seg.left[now]
			r = mid
		} else {
			now = seg.right[now]
			l = mid
		}
	}
	seg.data = append(seg.data, val)
	seg.left, seg.right = append(seg.left, 0), append(seg.right, 0)
	now = int32(len(seg.data) - 1)
	for i := len(path) - 2; i >= 0; i-- {
		if seg.left[path[i]] == path[i+1] {
			seg.data = append(seg.data, seg.op(seg.data[now], seg.data[seg.right[path[i]]]))
			seg.left, seg.right = append(seg.left, now), append(seg.right, seg.right[path[i]])
		} else {
			seg.data = append(seg.data, seg.op(seg.data[seg.left[path[i]]], seg.data[now]))
			seg.left, seg.right = append(seg.left, seg.left[path[i]]), append(seg.right, now)
		}
		now = int32(len(seg.data) - 1)
	}
	seg.root = append(seg.root, int32(len(seg.data)-1))
	return len(seg.root) - 1
}
func (seg *segment) query(ver, l, r int) Node {
	var qu func(pos int32, nodel, noder int) Node
	qu = func(pos int32, nodel, noder int) Node {
		switch {
		case l <= nodel && noder <= r:
			return seg.data[pos]
		case r <= nodel || noder <= l:
			return seg.e
		default:
			mid := (nodel + noder) / 2
			return seg.op(qu(seg.left[pos], nodel, mid), qu(seg.right[pos], mid, noder))
		}
	}
	return qu(seg.root[ver], 0, seg.n)
}

// max_right(l,f) -> f(op(a[l..r))) が真となる最大の r（半開区間）
// 前提：f(e)==true、かつ f は r について単調（真→偽に一度だけ変わる）
func (seg *segment) maxRight(ver, l int, f func(Node) bool) int {
	if l == seg.m {
		return l
	}
	var sm Node = seg.e
	var mr func(pos int32, nodel, noder int) int
	mr = func(pos int32, nodel, noder int) int {
		switch {
		case noder <= l:
			return -1
		case nodel >= l && f(seg.op(sm, seg.data[pos])):
			sm = seg.op(sm, seg.data[pos])
			return -1
		case noder-nodel == 1:
			return nodel
		default:
			mid := (nodel + noder) / 2
			p := mr(seg.left[pos], nodel, mid)
			if p != -1 {
				return p
			} else {
				return mr(seg.right[pos], mid, noder)
			}
		}
	}
	p := mr(seg.root[ver], 0, seg.n)
	if p == -1 {
		return seg.m
	}
	return p
}

// min_left(r,f) -> f(op(a[l..r))) が真となる最小の l（半開区間）
// 前提：f(e)==true、かつ f は l について単調
func (seg *segment) minLeft(ver, r int, f func(Node) bool) int {
	if r == 0 {
		return r
	}
	var sm Node = seg.e
	var ml func(pos int32, nodel, noder int) int
	ml = func(pos int32, nodel, noder int) int {
		switch {
		case nodel >= r:
			return -1
		case noder <= r && f(seg.op(seg.data[pos], sm)):
			sm = seg.op(seg.data[pos], sm)
			return -1
		case noder-nodel == 1:
			return noder
		default:
			mid := (nodel + noder) / 2
			p := ml(seg.right[pos], mid, noder)
			if p != -1 {
				return p
			} else {
				return ml(seg.left[pos], nodel, mid)
			}
		}
	}
	p := ml(seg.root[ver], 0, seg.n)
	if p == -1 {
		return 0
	}
	return p
}
