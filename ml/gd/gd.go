package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type GradientDescentModel struct {
	w float64
	b float64
}

func (m *GradientDescentModel) Fit(xs, ys []float64, a float64, epochs int) {
	for i := 0; i < epochs; i++ {
		m.learn(xs, ys, a)

		if i%400 == 0 {
			log.Printf("epoch: %d; loss: %f", i, m.mse(xs, ys))
		}
	}
}

func (m *GradientDescentModel) Predict(x float64) float64 {
	return m.w*x + m.b
}

func (m *GradientDescentModel) learn(xs, ys []float64, a float64) {
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

func (m *GradientDescentModel) mse(xs, ys []float64) float64 {
	n := len(xs)

	e := 0.0
	for i := 0; i < n; i++ {
		d := ys[i] - (m.w*xs[i] + m.b)
		e += d * d
	}
	e /= float64(n)

	return e
}

func main() {
	f, err := os.Open("advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	//goland:noinspection GoUnhandledErrorResult
	defer f.Close()

	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var spendings []float64
	var sales []float64
	for _, record := range records {
		spending, _ := strconv.ParseFloat(record[2], 64)
		spendings = append(spendings, spending)
		sale, _ := strconv.ParseFloat(record[4], 64)
		sales = append(sales, sale)
	}

	model := GradientDescentModel{}
	model.Fit(spendings, sales, 0.001, 15000)

	fmt.Println(model)
	fmt.Println(model.Predict(23))
}
