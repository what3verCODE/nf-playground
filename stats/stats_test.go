package stats

import (
	"errors"
	"testing"
)

func TestSum(t *testing.T) {
	for _, c := range []struct {
		in   []float64
		want float64
	}{{nil, 0}, {[]float64{1, 2, 3.5}, 6.5}, {[]float64{-1, 1}, 0}} {
		if got := Sum(c.in); got != c.want {
			t.Errorf("Sum(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestMax(t *testing.T) {
	if _, err := Max(nil); !errors.Is(err, ErrEmpty) {
		t.Fatalf("Max(nil): %v", err)
	}
	if got, _ := Max([]float64{-3, 7, 2}); got != 7 {
		t.Fatalf("Max = %v", got)
	}
}
