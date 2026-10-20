package main

//trie

const sigma int = 26
const base byte = 'a'

type trieTree struct {
	next, pass, end []int32
}

func f(b byte) int {
	return int(b - base)
}
func revF(x int) byte {
	return byte(x + int(base))
}
func newTrie(total int) trieTree {
	return trieTree{make([]int32, sigma, sigma*(total+1)), make([]int32, 1, total+1), make([]int32, 1, total+1)}
}
func (trie *trieTree) newNode() int {
	id := len(trie.pass)
	trie.next = append(trie.next, make([]int32, sigma)...)
	trie.pass, trie.end = append(trie.pass, 0), append(trie.end, 0)
	return id
}
func (trie *trieTree) insert(s string) {
	var v int = 0
	trie.pass[v]++
	for i := 0; i < len(s); i++ {
		c := int(trie.next[v*sigma+f(s[i])])
		if c == 0 {
			id := trie.newNode()
			trie.next[v*sigma+f(s[i])] = int32(id)
			c = id
		}
		v = c
		trie.pass[v]++
	}
	trie.end[v]++
}
func (trie *trieTree) find(s string) int {
	var v int = 0
	for i := 0; i < len(s); i++ {
		c := int(trie.next[v*sigma+f(s[i])])
		if c == 0 || trie.pass[c] == 0 {
			return -1
		}
		v = c
	}
	return v
}
func (trie *trieTree) count(s string) int32 {
	x := trie.find(s)
	if x == -1 {
		return 0
	}
	return trie.end[x]
}
func (trie *trieTree) countPrefix(s string) int32 {
	x := trie.find(s)
	if x == -1 {
		return 0
	}
	return trie.pass[x]
}
func (trie *trieTree) erase(s string) bool {
	if trie.count(s) == 0 {
		return false
	}
	var v int = 0
	trie.pass[v]--
	for i := 0; i < len(s); i++ {
		c := int(trie.next[v*sigma+f(s[i])])
		trie.pass[c]--
		v = c
	}
	trie.end[v]--
	return true
}
func (trie *trieTree) LCP(s string) int {
	v := 0
	for i := 0; i < len(s); i++ {
		c := int(trie.next[v*sigma+f(s[i])])
		if c == 0 || trie.pass[c] == 0 {
			return i
		}
		v = c
	}
	return len(s)
}
func (trie *trieTree) Kth(k int) (string, bool) {
	if k < 0 || trie.size() <= k {
		return "", false
	}
	var buf []byte
	var v int = 0
	for {
		e := int(trie.end[v])
		if k < e {
			return string(buf), true
		}
		k -= e
		suc := false
		for i := 0; i < sigma; i++ {
			c := int(trie.next[v*sigma+i])
			cnt := int(trie.pass[c])
			if c == 0 || cnt == 0 {
				continue
			}
			if k < cnt {
				buf = append(buf, revF(i))
				v = c
				suc = true
				break
			}
			k -= cnt
		}
		if !suc {
			return "", false
		}
	}
}
func (trie *trieTree) size() int {
	return int(trie.pass[0])
}
