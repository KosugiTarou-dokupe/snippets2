func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
func lcm(a, b int64) int64 {
	return a / gcd(a, b) * b
}
