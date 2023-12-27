package game

import (
	"github.com/google/go-cmp/cmp"
	"math/rand"
	"testing"
)

func TestNew(t *testing.T) {
	got := New(rand.New(rand.NewSource(42)))
	want := &State{
		Status: StatusPlaying,
		Board: [4][4]int{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 2},
			{0, 0, 4, 0},
		},
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Error(diff)
	}
}

type dummyRandSource struct{}

func (s *dummyRandSource) Int63() int64 {
	return 0 // This gives us always first emtpy cell filled with 4.
}

func (s *dummyRandSource) Seed(seed int64) {}

func TestState_Move(t *testing.T) {
	for _, tt := range []struct {
		name      string
		state     *State
		direction Direction
		want      *State
	}{
		{
			name: "simple move to left",
			state: &State{
				Status: StatusPlaying,
				Board: [4][4]int{
					{0, 0, 0, 0},
					{0, 2, 0, 4},
					{4, 0, 2, 0},
					{0, 0, 0, 0},
				},
			},
			direction: DirectionLeft,
			want: &State{
				Status: StatusPlaying,
				Board: [4][4]int{
					{4, 0, 0, 0},
					{2, 4, 0, 0},
					{4, 2, 0, 0},
					{0, 0, 0, 0},
				},
			},
		},
		{
			name: "simple move to top",
			state: &State{
				Status: StatusPlaying,
				Board: [4][4]int{
					{0, 0, 2, 0},
					{0, 4, 0, 0},
					{0, 0, 4, 0},
					{0, 2, 0, 0},
				},
			},
			direction: DirectionUp,
			want: &State{
				Status: StatusPlaying,
				Board: [4][4]int{
					{4, 4, 2, 0},
					{0, 2, 4, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
				},
			},
		},
		{
			name: "simple move to right",
			state: &State{
				Status: StatusPlaying,
				Board: [4][4]int{
					{0, 0, 0, 0},
					{0, 4, 0, 2},
					{2, 0, 4, 0},
					{0, 0, 0, 0},
				},
			},
			direction: DirectionRight,
			want: &State{
				Status: StatusPlaying,
				Board: [4][4]int{
					{4, 0, 0, 0},
					{0, 0, 4, 2},
					{0, 0, 2, 4},
					{0, 0, 0, 0},
				},
			},
		},
		{
			name: "simple move to down",
			state: &State{
				Status: StatusPlaying,
				Board: [4][4]int{
					{0, 0, 2, 0},
					{0, 4, 0, 0},
					{0, 0, 4, 0},
					{0, 2, 0, 0},
				},
			},
			direction: DirectionDown,
			want: &State{
				Status: StatusPlaying,
				Board: [4][4]int{
					{4, 0, 0, 0},
					{0, 0, 0, 0},
					{0, 4, 2, 0},
					{0, 2, 4, 0},
				},
			},
		},
		{
			name: "collapse to left",
			state: &State{
				Status: StatusPlaying,
				Score:  4,
				Board: [4][4]int{
					{2, 2, 4, 4},
					{2, 4, 2, 4},
					{0, 4, 0, 4},
					{0, 0, 0, 0},
				},
			},
			direction: DirectionLeft,
			want: &State{
				Status: StatusPlaying,
				Score:  24,
				Board: [4][4]int{
					{4, 8, 4, 0},
					{2, 4, 2, 4},
					{8, 0, 0, 0},
					{0, 0, 0, 0},
				},
			},
		},
		{
			name: "collapse to right",
			state: &State{
				Status: StatusPlaying,
				Score:  4,
				Board: [4][4]int{
					{2, 2, 4, 4},
					{2, 4, 2, 4},
					{0, 4, 0, 4},
					{0, 0, 0, 0},
				},
			},
			direction: DirectionRight,
			want: &State{
				Status: StatusPlaying,
				Score:  24,
				Board: [4][4]int{
					{4, 0, 4, 8},
					{2, 4, 2, 4},
					{0, 0, 0, 8},
					{0, 0, 0, 0},
				},
			},
		},
		{
			name: "collapse to top",
			state: &State{
				Status: StatusPlaying,
				Score:  4,
				Board: [4][4]int{
					{2, 2, 4, 4},
					{2, 4, 2, 4},
					{0, 4, 0, 4},
					{0, 0, 0, 0},
				},
			},
			direction: DirectionUp,
			want: &State{
				Status: StatusPlaying,
				Score:  24,
				Board: [4][4]int{
					{4, 2, 4, 8},
					{4, 8, 2, 4},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
				},
			},
		},
		{
			name: "collapse to bottom",
			state: &State{
				Status: StatusPlaying,
				Score:  4,
				Board: [4][4]int{
					{2, 2, 4, 4},
					{2, 4, 2, 4},
					{0, 4, 0, 4},
					{0, 0, 0, 0},
				},
			},
			direction: DirectionDown,
			want: &State{
				Status: StatusPlaying,
				Score:  24,
				Board: [4][4]int{
					{4, 0, 0, 0},
					{0, 0, 0, 0},
					{0, 2, 4, 4},
					{4, 8, 2, 8},
				},
			},
		},
		{
			name: "game over",
			state: &State{
				Status: StatusPlaying,
				Score:  4,
				Board: [4][4]int{
					{2, 4, 2, 4},
					{4, 2, 4, 2},
					{0, 2, 4, 2},
					{4, 2, 4, 2},
				},
			},
			direction: DirectionLeft,
			want: &State{
				Status: StatusOver,
				Score:  4,
				Board: [4][4]int{
					{2, 4, 2, 4},
					{4, 2, 4, 2},
					{2, 4, 2, 4},
					{4, 2, 4, 2},
				},
			},
		},
		{
			name: "full but can move horizontally",
			state: &State{
				Status: StatusPlaying,
				Score:  4,
				Board: [4][4]int{
					{2, 4, 2, 4},
					{4, 2, 4, 2},
					{0, 2, 8, 4},
					{4, 2, 16, 2},
				},
			},
			direction: DirectionLeft,
			want: &State{
				Status: StatusPlaying,
				Score:  4,
				Board: [4][4]int{
					{2, 4, 2, 4},
					{4, 2, 4, 2},
					{2, 8, 4, 4},
					{4, 2, 16, 2},
				},
			},
		},
		{
			name: "full but can move vertically",
			state: &State{
				Status: StatusPlaying,
				Score:  4,
				Board: [4][4]int{
					{2, 4, 2, 4},
					{4, 2, 4, 2},
					{0, 2, 8, 16},
					{4, 2, 32, 4},
				},
			},
			direction: DirectionLeft,
			want: &State{
				Status: StatusPlaying,
				Score:  4,
				Board: [4][4]int{
					{2, 4, 2, 4},
					{4, 2, 4, 2},
					{2, 8, 16, 4},
					{4, 2, 32, 4},
				},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.state.Move(tt.direction, rand.New(&dummyRandSource{}))
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Error(diff)
			}
		})
	}
}
