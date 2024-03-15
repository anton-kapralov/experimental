package calculus

import (
	"math"
	"testing"

	"gonum.org/v1/gonum/stat/distuv"
)

const floatTolerance = 1.0e-9

func TestDefiniteIntegral(t *testing.T) {
	dist := &distuv.Normal{
		Mu:    10,
		Sigma: 1,
	}

	for _, tt := range []struct {
		name string
		f    func(float64) float64
		want float64
	}{
		{
			name: "pdf",
			f:    dist.Prob,
			want: 1.0,
		},
		{
			name: "expected value",
			f:    func(x float64) float64 { return x * dist.Prob(x) },
			want: 10.0,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := DefiniteIntegral(tt.f, 0, 20)

			if math.Abs(got-tt.want) > floatTolerance {
				t.Errorf("DefiniteIntegral(..) = %g; want %g", got, tt.want)
			}
		})
	}
}
