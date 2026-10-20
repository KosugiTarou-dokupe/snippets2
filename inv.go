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
