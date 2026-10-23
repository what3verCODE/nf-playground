// Package stats has small helpers over slices of numbers.
package stats

import "errors"

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

// Min is the smallest number, or ErrEmpty.
func Min(xs []float64) (float64, error) {
	if len(xs) == 0 {
		return 0, ErrEmpty
	}
	m := xs[0]
	for _, x := range xs[1:] {
		if x < m {
			m = x
		}
	}
	return m, nil
}
