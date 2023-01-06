package gradientdescent

import "log"

type Model struct {
	w float64
	b float64
}

func NewModel(w, b float64) *Model {
	return &Model{
		w: w,
		b: b,
	}
}

func (m *Model) Fit(xs, ys []float64, a float64, epochs int) {
	for i := 0; i < epochs; i++ {
		m.learn(xs, ys, a)

		if i%400 == 0 {
			log.Printf("epoch: %d; loss: %f", i, m.mse(xs, ys))
		}
	}
}

func (m *Model) Predict(x float64) float64 {
	return m.w*x + m.b
}

func (m *Model) learn(xs, ys []float64, a float64) {
	n := len(xs)

	dw := 0.0
	db := 0.0
	for i := 0; i < n; i++ {
		d := ys[i] - (m.w*xs[i] + m.b)
		dw += -2 * xs[i] * d
		db += -2 * d
	}
	dw /= float64(n)
	db /= float64(n)

	m.w -= a * dw
	m.b -= a * db
}

func (m *Model) mse(xs, ys []float64) float64 {
	n := len(xs)

	e := 0.0
	for i := 0; i < n; i++ {
		d := ys[i] - (m.w*xs[i] + m.b)
		e += d * d
	}
	e /= float64(n)

	return e
}
