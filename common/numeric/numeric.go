package numeric

// Gcd2 returns the greatest common divisor of two integers
func Gcd2(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Gcd returns the greatest common divisor of a list of integers
func Gcd(values []int) int {
	g := values[0]
	for _, v := range values[1:] {
		g = Gcd2(g, v)
	}
	return g
}

// Lcm2 returns the least common multiple of two integers
func Lcm2(a, b int) int {
	return a * b / Gcd2(a, b)
}

// Lcm returns the least common multiple of a list of integers
func Lcm(values []int) int {
	l := values[0]
	for _, v := range values[1:] {
		l = Lcm2(l, v)
	}
	return l
}
