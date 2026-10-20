type node struct {
	x int64
}
type lazytp struct {
	x int64
}
type lazyseg struct {
	n, m        int
	data        []node
	lazy        []lazytp
	e           node
	id          lazytp
	op          func(a, b node) node
	composition func(a, b lazytp) lazytp
	mapping     func(a lazytp, b node, len int) node
}

func newlazyseg(arr []node, e node, id lazytp, op func(a, b node) node, composition func(a, b lazytp) lazytp, mapping func(a lazytp, b node, len int) node) *lazyseg {
	var n int = 1
	for n < len(arr) {
		n *= 2
	}
	data, lazy := make([]node, 2*n), make([]lazytp, 2*n)
	for i := 0; i < 2*n; i++ {
		data[i], lazy[i] = e, id
	}
	for i := 0; i < len(arr); i++ {
		data[n+i] = arr[i]
	}
	for i := n - 1; i >= 0; i-- {
		data[i] = op(data[2*i], data[2*i+1])
	}
	return &lazyseg{n, len(arr), data, lazy, e, id, op, composition, mapping}
}
func (lazy *lazyseg) push(k, len int) {
	if lazy.lazy[k] == lazy.id {
		return
	}
	lazy.data[2*k] = lazy.mapping(lazy.lazy[k], lazy.data[2*k], len/2)
	lazy.lazy[2*k] = lazy.composition(lazy.lazy[k], lazy.lazy[2*k])
	lazy.data[2*k+1] = lazy.mapping(lazy.lazy[k], lazy.data[2*k+1], len/2)
	lazy.lazy[2*k+1] = lazy.composition(lazy.lazy[k], lazy.lazy[2*k+1])
	lazy.lazy[k] = lazy.id
}
func (lazy *lazyseg) update(pos, i, nodel, noder int, val lazytp) {
	if noder-nodel == 1 {
		lazy.data[pos] = lazy.mapping(val, lazy.data[pos], 1)
		lazy.lazy[pos] = lazy.composition(val, lazy.lazy[pos])
		return
	}
	lazy.push(pos, noder-nodel)
	mid := (noder + nodel) / 2
	if i < mid {
		lazy.update(pos*2, i, nodel, mid, val)
	} else {
		lazy.update(pos*2+1, i, mid, noder, val)
	}
	lazy.data[pos] = lazy.op(lazy.data[pos*2], lazy.data[pos*2+1])
}
func (lazy *lazyseg) lenUpd(pos, l, r, nodel, noder int, val lazytp) {
	if r <= nodel || noder <= l {
		return
	}
	if l <= nodel && noder <= r {
		lazy.data[pos] = lazy.mapping(val, lazy.data[pos], noder-nodel)
		lazy.lazy[pos] = lazy.composition(val, lazy.lazy[pos])
	} else {
		lazy.push(pos, noder-nodel)
		mid := (noder + nodel) / 2
		lazy.lenUpd(pos*2, l, r, nodel, mid, val)
		lazy.lenUpd(pos*2+1, l, r, mid, noder, val)
		lazy.data[pos] = lazy.op(lazy.data[pos*2], lazy.data[pos*2+1])
	}
}
func (lazy *lazyseg) Query(pos, l, r, nodel, noder int) node {
	if r <= nodel || noder <= l {
		return lazy.e
	}
	if l <= nodel && noder <= r {
		return lazy.data[pos]
	}
	lazy.push(pos, noder-nodel)
	mid := (noder + nodel) / 2
	return lazy.op(lazy.Query(pos*2, l, r, nodel, mid), lazy.Query(pos*2+1, l, r, mid, noder))
}
func (lazy *lazyseg) mr(pos, nodel, noder, l int, sm *node, f func(node) bool) int {
	if noder <= l {
		return -1
	}
	if l <= nodel {
		if f(lazy.op(*sm, lazy.data[pos])) {
			*sm = lazy.op(*sm, lazy.data[pos])
			return -1
		}
		if noder-nodel == 1 {
			return nodel
		}
	}
	lazy.push(pos, noder-nodel)
	mid := (noder + nodel) / 2
	ret := lazy.mr(pos*2, nodel, mid, l, sm, f)
	if ret != -1 {
		return ret
	}
	return lazy.mr(pos*2+1, mid, noder, l, sm, f)
}
func (lazy *lazyseg) ml(pos, nodel, noder, r int, sm *node, f func(node) bool) int {
	if r <= nodel {
		return -1
	}
	if noder <= r {
		if f(lazy.op(lazy.data[pos], *sm)) {
			*sm = lazy.op(lazy.data[pos], *sm)
			return -1
		}
		if noder-nodel == 1 {
			return noder
		}
	}
	lazy.push(pos, noder-nodel)
	mid := (noder + nodel) / 2
	ret := lazy.ml(pos*2+1, mid, noder, r, sm, f)
	if ret != -1 {
		return ret
	}
	return lazy.ml(pos*2, nodel, mid, r, sm, f)
}
func (lazy *lazyseg) maxRight(l int, f func(node) bool) int {
	sm := lazy.e
	ret := lazy.mr(1, 0, lazy.n, l, &sm, f)
	if ret == -1 {
		return lazy.m
	}
	return ret
}
func (lazy *lazyseg) minLeft(r int, f func(node) bool) int {
	sm := lazy.e
	ret := lazy.ml(1, 0, lazy.n, r, &sm, f)
	if ret == -1 {
		return 0
	}
	return ret
}
