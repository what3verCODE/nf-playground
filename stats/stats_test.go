package stats

import (
	"errors"
	"slices"
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

func TestMean(t *testing.T) {
	if _, err := Mean(nil); !errors.Is(err, ErrEmpty) {
		t.Fatalf("Mean(nil): %v", err)
	}
	if _, err := Mean([]float64{}); !errors.Is(err, ErrEmpty) {
		t.Fatalf("Mean([]): %v", err)
	}
	for _, c := range []struct {
		in   []float64
		want float64
	}{
		{[]float64{4}, 4},
		{[]float64{1, 2, 3, 4}, 2.5},
		{[]float64{1, 2, 6}, 3},
		{[]float64{-4, -2, 0}, -2},
	} {
		if got, err := Mean(c.in); err != nil || got != c.want {
			t.Errorf("Mean(%v) = %v, %v; want %v", c.in, got, err, c.want)
		}
	}
}

func TestMedian(t *testing.T) {
	if _, err := Median(nil); !errors.Is(err, ErrEmpty) {
		t.Fatalf("Median(nil): %v", err)
	}
	if _, err := Median([]float64{}); !errors.Is(err, ErrEmpty) {
		t.Fatalf("Median([]): %v", err)
	}
	for _, c := range []struct {
		in   []float64
		want float64
	}{
		{[]float64{4}, 4},
		{[]float64{3, 1, 2}, 2},
		{[]float64{4, 1, 3, 2}, 2.5},
		{[]float64{-5, -1, -3}, -3},
		{[]float64{-4, 2, -2, 6}, 0},
	} {
		orig := append([]float64(nil), c.in...)
		if got, err := Median(c.in); err != nil || got != c.want {
			t.Errorf("Median(%v) = %v, %v; want %v", orig, got, err, c.want)
		}
		if !slices.Equal(c.in, orig) {
			t.Errorf("Median reordered input: %v, was %v", c.in, orig)
		}
	}
}
