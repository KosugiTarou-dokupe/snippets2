package main

//nCr

func newFacts(n int, mod int64, facts []int64) []int64 {
	if len(facts) >= n+1 {
		return facts
	}
	if len(facts) == 0 {
		facts = []int64{1}
	}
	for i := len(facts); i < n+1; i++ {
		facts = append(facts, facts[i-1]*int64(i)%mod)
	}
	return facts
}
func newInvFacts(n int, mod int64, facts, invFacts []int64) []int64 {
	if len(invFacts) >= n+1 {
		return invFacts
	}
	l := len(invFacts)
	invFacts = append(invFacts, make([]int64, n+1-l)...)
	invFacts[n] = modinv(facts[n], mod)
	for i := n - 1; i >= l; i-- {
		invFacts[i] = invFacts[i+1] * int64(i+1) % mod
	}
	return invFacts
}
func modinv(n, q int64) int64 {
	n %= q
	n = (n + q) % q
	if n == 0 {
		panic("逆元が存在しない")
	}
	return modpow(n, q-2, q)
}
func modpow(n, a, m int64) int64 {
	ret := int64(1)
	for a > 0 {
		if a&1 == 1 {
			ret = ret * n % m
		}
		n = n * n % m
		a >>= 1
	}
	return ret
}
func nCr(n, r, mod int64, facts, invFacts []int64) int64 {
	return ((facts[n] * invFacts[r]) % mod * invFacts[n-r]) % mod
}
