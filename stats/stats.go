// Package stats has small helpers over slices of numbers.
package stats

import (
	"errors"
	"slices"
)

// ErrEmpty is returned for an empty input where a value is undefined.
var ErrEmpty = errors.New("stats: empty input")

// Sum adds the numbers; 0 for none.
func Sum(xs []float64) float64 {
	var s float64
	for _, x := range xs {
		s += x
	}
	return s
}

// Max is the largest number, or ErrEmpty.
func Max(xs []float64) (float64, error) {
	if len(xs) == 0 {
		return 0, ErrEmpty
	}
	m := xs[0]
	for _, x := range xs[1:] {
		if x > m {
			m = x
		}
	}
	return m, nil
}

// Mean is the arithmetic mean, or ErrEmpty.
func Mean(xs []float64) (float64, error) {
	if len(xs) == 0 {
		return 0, ErrEmpty
	}
	return Sum(xs) / float64(len(xs)), nil
}

// Median is the middle number (mean of the two middle ones for an even count), or ErrEmpty.
// It sorts a copy and leaves xs as it was.
func Median(xs []float64) (float64, error) {
	if len(xs) == 0 {
		return 0, ErrEmpty
	}
	s := slices.Sorted(slices.Values(xs))
	n := len(s)
	if n%2 == 1 {
		return s[n/2], nil
	}
	return (s[n/2-1] + s[n/2]) / 2, nil
}
