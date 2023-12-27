package rest

import (
	"github.com/google/go-cmp/cmp"
	"math/rand"
	"testing"
)

func TestNewGameKey(t *testing.T) {
	got, want := newGameKey(rand.New(rand.NewSource(42)), 16), "3BsYdLdigjpvPgKz"
	if diff := cmp.Diff(got, want); diff != "" {
		t.Error(diff)
	}
}
