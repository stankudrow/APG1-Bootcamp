package ex00_tests

import (
	sts "ex00/anscombe"
	"math"
	"testing"
)

var TOL float64 = 1e-6

func TestGetMean(t *testing.T) {
	tests := map[string]struct {
		numbers  []float64
		expected float64
		is_err   bool
	}{
		"{}":              {make([]float64, 0), math.NaN(), true},
		"{-1, 0, 1}":      {[]float64{-1, 0, 1}, 0, false},
		"{1.1, 2.2, 3.3}": {[]float64{1.1, 2.2, 3.3}, 2.2, false},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := sts.GetMean(test.numbers)
			if test.is_err {
				if err == nil {
					t.Errorf("no error has been thrown")
				}
			} else {
				if math.Abs(result-test.expected) >= TOL {
					t.Errorf("result is %f, expected is %f", result, test.expected)
				}
			}
		})
	}
}

func TestGetMedian(t *testing.T) {
	tests := map[string]struct {
		numbers  []float64
		expected float64
		is_err   bool
	}{
		"{}":           {make([]float64, 0), math.NaN(), true},
		"{1}":          {[]float64{1}, 1, false},
		"{1, 2}":       {[]float64{1, 2}, 1.5, false},
		"{1, 2, 3}":    {[]float64{1, 2, 3}, 2, false},
		"{1, 2, 3, 4}": {[]float64{1, 2, 3, 4}, 2.5, false},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := sts.GetMedian(test.numbers)
			if test.is_err {
				if err == nil {
					t.Errorf("no error has been thrown")
				}
			} else {
				if math.Abs(result-test.expected) >= TOL {
					t.Errorf("result is %f, expected is %f", result, test.expected)
				}
			}
		})
	}
}

func TestGetMode(t *testing.T) {
	tests := map[string]struct {
		numbers  []float64
		expected float64
		is_err   bool
	}{
		"{}":                 {make([]float64, 0), math.NaN(), true},
		"{1}":                {[]float64{1}, 1, false},
		"{0, 1, 1, 2, 2, 2}": {[]float64{0, 1, 1, 2, 2, 2}, 2, false},
		"{0, 1, 1, 3, 3}":    {[]float64{0, 1, 1, 3, 3}, 1, false},
		"{0, 1, 2}":          {[]float64{0, 1, 2}, 0, false},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := sts.GetMode(test.numbers)
			if test.is_err {
				if err == nil {
					t.Errorf("no error has been thrown")
				}
			} else {
				if math.Abs(result-test.expected) >= TOL {
					t.Errorf("result is %f, expected is %f", result, test.expected)
				}
			}
		})
	}
}

func TestGetPopulationStandardDeviation(t *testing.T) {
	tests := map[string]struct {
		numbers  []float64
		expected float64
		is_err   bool
	}{
		"{}":            {make([]float64, 0), math.NaN(), true},
		"{1}":           {[]float64{1}, 0, false},
		"{1, 3}":        {[]float64{1, 3}, 1, false},
		"{-2, 0, 2, 4}": {[]float64{-2, 0, 2, 4}, math.Sqrt(5), false},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := sts.GetPopulationStandardDeviation(test.numbers)
			if test.is_err {
				if err == nil {
					t.Errorf("no error has been thrown")
				}
			} else {
				if math.Abs(result-test.expected) >= TOL {
					t.Errorf("result is %f, expected is %f", result, test.expected)
				}
			}
		})
	}
}

func TestGetSampleStandardDeviation(t *testing.T) {
	tests := map[string]struct {
		numbers  []float64
		expected float64
		is_err   bool
	}{
		"{}":            {make([]float64, 0), math.NaN(), true},
		"{1}":           {[]float64{1}, math.NaN(), true},
		"{1, 3}":        {[]float64{1, 3}, math.Sqrt(2), false},
		"{-2, 0, 2, 4}": {[]float64{-2, 0, 2, 4}, math.Sqrt(20.0 / 3.0), false},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := sts.GetSampleStandardDeviation(test.numbers)
			if test.is_err {
				if err == nil {
					t.Errorf("no error has been thrown")
				}
			} else {
				if math.Abs(result-test.expected) >= TOL {
					t.Errorf("result is %f, expected is %f", result, test.expected)
				}
			}
		})
	}
}
