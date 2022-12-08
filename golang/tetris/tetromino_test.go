package tetris

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestPoint2DRotate(t *testing.T) {
	for _, tt := range []struct {
		name      string
		point     point2D
		center    point2D
		direction rotationDirection
		want      point2D
	}{
		{
			name:      "clockwise around (0,0)",
			point:     point2D{5, 3},
			center:    point2D{},
			direction: clockwise,
			want:      point2D{3, -5},
		},
		{
			name:      "counter-clockwise around (0,0)",
			point:     point2D{5, 3},
			center:    point2D{},
			direction: counterClockwise,
			want:      point2D{-3, 5},
		},
		{
			name:      "clockwise around (1,1)",
			point:     point2D{0, 1},
			center:    point2D{1, 1},
			direction: clockwise,
			want:      point2D{1, 2},
		},
		{
			name:      "counter-clockwise around (1,1)",
			point:     point2D{0, 1},
			center:    point2D{1, 1},
			direction: counterClockwise,
			want:      point2D{1, 0},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			tt.point.rotate(tt.center, tt.direction)
			if diff := cmp.Diff(tt.point, tt.want); diff != "" {
				t.Error(diff)
			}
		})
	}
}
