func mulmod(a, b, mod uint64) uint64 {
	a %= mod
	var ret uint64 = 0
	for b > 0 {
		if b&1 == 1 {
			ret = (ret + a) % mod
		}
		a = (a * 2) % mod
		b = b >> 1
	}
	return ret
}
