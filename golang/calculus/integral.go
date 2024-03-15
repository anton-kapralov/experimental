package calculus

const precision = 1000

func DefiniteIntegral(f func(float64) float64, a, b float64) float64 {
	x := a
	dx := (b - a) / float64(precision)
	s := 0.0
	for i := 0; i < precision; i++ {
		s += f(x) * dx
		x += dx
	}
	return s
}
