import "cmp"

type mnode[T any] struct {
	key          T
	cnt, sz, sum int
	pri          uint32
	l, r         *mnode[T]
}
type mset[T any] struct {
	root *mnode[T]
	seed uint32
	cmp  func(a, b T) int
	val  func(x T) int
}

func newmset[T any](cmp func(a, b T) int, val func(x T) int) *mset[T] {
	return &mset[T]{nil, 2463534242, cmp, val}
}
func newintmset() *mset[int] {
	return newmset(cmp.Compare, func(x int) int { return x })
}
func (m *mset[T]) rnd() uint32 {
	m.seed = m.seed ^ (m.seed << 13)
	m.seed = m.seed ^ (m.seed >> 17)
	m.seed = m.seed ^ (m.seed << 5)
	return m.seed
}
func (m *mset[T]) msz(n *mnode[T]) int {
	if n == nil {
		return 0
	}
	return n.sz
}
func (m *mset[T]) msum(n *mnode[T]) int {
	if n == nil {
		return 0
	}
	return n.sum
}
func (m *mset[T]) upd(n *mnode[T]) {
	n.sz = n.cnt + m.msz(n.l) + m.msz(n.r)
	n.sum = n.cnt*m.val(n.key) + m.msum(n.l) + m.msum(n.r)
}
func (m *mset[T]) rotR(n *mnode[T]) *mnode[T] {
	l := n.l
	n.l = l.r
	l.r = n
	m.upd(n)
	m.upd(l)
	return l
}
func (m *mset[T]) rotL(n *mnode[T]) *mnode[T] {
	r := n.r
	n.r = r.l
	r.l = n
	m.upd(n)
	m.upd(r)
	return r
}
func (m *mset[T]) ins(n *mnode[T], key T) *mnode[T] {
	if n == nil {
		return &mnode[T]{key, 1, 1, m.val(key), m.rnd(), nil, nil}
	}
	switch m.cmp(key, n.key) {
	case -1:
		n.l = m.ins(n.l, key)
		m.upd(n)
		if n.pri < n.l.pri {
			n = m.rotR(n)
		}
	case 1:
		n.r = m.ins(n.r, key)
		m.upd(n)
		if n.pri < n.r.pri {
			n = m.rotL(n)
		}
	default:
		n.cnt++
		m.upd(n)
	}
	return n
}
func (m *mset[T]) del(n *mnode[T]) *mnode[T] {
	if n.l == nil {
		return n.r
	}
	if n.r == nil {
		return n.l
	}
	if n.l.pri > n.r.pri {
		n = m.rotR(n)
		n.r = m.del(n.r)
	} else {
		n = m.rotL(n)
		n.l = m.del(n.l)
	}
	m.upd(n)
	return n
}
func (m *mset[T]) era(n *mnode[T], key T, all bool) *mnode[T] {
	if n == nil {
		return nil
	}
	switch m.cmp(key, n.key) {
	case -1:
		n.l = m.era(n.l, key, all)
		m.upd(n)
		return n
	case 0:
		if !all && n.cnt > 1 {
			n.cnt--
			m.upd(n)
			return n
		}
		return m.del(n)
	default:
		n.r = m.era(n.r, key, all)
		m.upd(n)
		return n
	}
}
func (m *mset[T]) insert(key T) {
	m.root = m.ins(m.root, key)
}
func (m *mset[T]) erase(key T) bool {
	if m.count(key) == 0 {
		return false
	}
	m.root = m.era(m.root, key, false)
	return true
}
func (m *mset[T]) eraseAll(key T) int {
	c := m.count(key)
	if c > 0 {
		m.root = m.era(m.root, key, true)
	}
	return c
}
func (m *mset[T]) eraseKth(k int) (T, bool) {
	v, ok := m.Kth(k)
	if ok {
		m.root = m.era(m.root, v, false)
	}
	return v, ok

}
func (m *mset[T]) popFirst() (T, bool) {
	return m.eraseKth(0)
}
func (m *mset[T]) popLast() (T, bool) {
	return m.eraseKth(m.size() - 1)
}
func (m *mset[T]) size() int {
	return m.msz(m.root)
}
func (m *mset[T]) count(key T) int {
	n := m.root
	for n != nil {
		switch m.cmp(key, n.key) {
		case -1:
			n = n.l
		case 0:
			return n.cnt
		default:
			n = n.r
		}
	}
	return 0
}
func (m *mset[T]) countBefore(x T) int {
	res, n := 0, m.root
	for n != nil {
		if m.cmp(n.key, x) < 0 {
			res += m.msz(n.l) + n.cnt
			n = n.r
		} else {
			n = n.l
		}
	}
	return res
}
func (m *mset[T]) countNotAfter(x T) int {
	return m.countBefore(x) + m.count(x)
}
func (m *mset[T]) countAfter(x T) int {
	return m.size() - m.countNotAfter(x)
}
func (m *mset[T]) countNotBefore(x T) int {
	return m.size() - m.countBefore(x)
}
func (m *mset[T]) countBetween(l, r T) int {
	if m.cmp(l, r) > 0 {
		return 0
	}
	return m.countNotAfter(r) - m.countBefore(l)
}
func (m *mset[T]) sumAll() int {
	return m.msum(m.root)
}
func (m *mset[T]) sumBefore(x T) int {
	res, n := 0, m.root
	for n != nil {
		if m.cmp(n.key, x) < 0 {
			res += m.msum(n.l) + m.val(n.key)*n.cnt
			n = n.r
		} else {
			n = n.l
		}
	}
	return res
}
func (m *mset[T]) sumNotAfter(x T) int {
	return m.sumBefore(x) + m.val(x)*m.count(x)
}
func (m *mset[T]) sumAfter(x T) int {
	return m.sumAll() - m.sumNotAfter(x)
}
func (m *mset[T]) sumNotBefore(x T) int {
	return m.sumAll() - m.sumBefore(x)
}
func (m *mset[T]) sumBetween(l, r T) int {
	if m.cmp(l, r) > 0 {
		return 0
	}
	return m.sumNotAfter(r) - m.sumBefore(l)
}
func (m *mset[T]) sumFirstK(k int) int {
	if k < 0 {
		return 0
	}
	res, n := 0, m.root
	for n != nil {
		if k >= m.msz(n) {
			res += m.msum(n)
			break
		}
		ls := m.msz(n.l)
		if k <= ls {
			n = n.l
			continue
		}
		res += m.msum(n.l)
		k -= ls
		if k <= n.cnt {
			res += m.val(n.key) * k
			break
		}
		res += m.val(n.key) * n.cnt
		k -= n.cnt
		n = n.r
	}
	return res
}
func (m *mset[T]) sumLastK(k int) int {
	if k < 0 {
		return 0
	}
	if k >= m.size() {
		return m.sumAll()
	}
	return m.sumAll() - m.sumFirstK(m.size()-k)
}
func (m *mset[T]) Kth(k int) (T, bool) {
	if k < 0 || k >= m.size() {
		var zero T
		return zero, false
	}
	n := m.root
	for {
		ls := m.msz(n.l)
		if k < ls {
			n = n.l
			continue
		}
		if k < ls+n.cnt {
			return n.key, true
		}
		k -= ls + n.cnt
		n = n.r
	}
}
func (m *mset[T]) KthLast(k int) (T, bool) {
	return m.Kth(m.size() - k - 1)
}
func (m *mset[T]) first() (T, bool) {
	return m.Kth(0)
}
func (m *mset[T]) last() (T, bool) {
	return m.Kth(m.size() - 1)
}
func (m *mset[T]) firstNotBefore(x T) (T, bool) {
	ok, n := false, m.root
	var res T
	for n != nil {
		if m.cmp(n.key, x) >= 0 {
			res = n.key
			ok = true
			n = n.l
		} else {
			n = n.r
		}
	}
	return res, ok
}
func (m *mset[T]) firstAfter(x T) (T, bool) {
	ok, n := false, m.root
	var res T
	for n != nil {
		if m.cmp(n.key, x) > 0 {
			res = n.key
			ok = true
			n = n.l
		} else {
			n = n.r
		}
	}
	return res, ok
}
func (m *mset[T]) lastNotAfter(x T) (T, bool) {
	ok, n := false, m.root
	var res T
	for n != nil {
		if m.cmp(x, n.key) >= 0 {
			res = n.key
			ok = true
			n = n.r
		} else {
			n = n.l
		}
	}
	return res, ok
}
func (m *mset[T]) lastBefore(x T) (T, bool) {
	ok, n := false, m.root
	var res T
	for n != nil {
		if m.cmp(x, n.key) > 0 {
			res = n.key
			ok = true
			n = n.r
		} else {
			n = n.l
		}
	}
	return res, ok
}
func (m *mset[T]) toSlice() []T {
	var ret []T = make([]T, 0, m.size())
	var f func(*mnode[T])
	f = func(n *mnode[T]) {
		if n == nil {
			return
		}
		f(n.l)
		for i := 0; i < n.cnt; i++ {
			ret = append(ret, n.key)
		}
		f(n.r)
	}
	f(m.root)
	return ret
}
