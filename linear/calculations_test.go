package linear

import (
	"testing"
)

func TestLinearRegression(t *testing.T) {
	tests := []struct {
		values    []float64
		expectedM float64
		expectedB float64
	}{
		{
			values:    []float64{189, 113, 121, 114, 145, 110},
			expectedM: -8.742857142857142,
			expectedB: 153.85714285714286,
		},
		{
			values:    []float64{189.8, 113.6, 121.7, 114.9, 145.1, 110},
			expectedM: -8.894285714285727,
			expectedB: 154.752380952381,
		},
		{
			values:    []float64{-189, 113, 121, -114, -145, 110},
			expectedM: 13.885714285714286,
			expectedB: -52.047619047619044,
		},
	}

	for _, test := range tests {
		m, b := LinearRegression(test.values)
		if m != test.expectedM || b != test.expectedB {
			t.Errorf("LinearRegression(%v) = (%.6f, %.6f), want (%.6f, %.6f)", test.values, m, b, test.expectedM, test.expectedB)
		}
	}
}

func TestCalculatePearson(t *testing.T) {
	tests := []struct {
		values    []float64
		expectedM float64
	}{
		{
			values:    []float64{189.8, 113.6, 121.7, 114.9, 145.1, 110},
			expectedM: -0.5408937201329931,
		},
		{
			values:    []float64{189, 113, 121, 114, 145, 110},
			expectedM: -0.5330331011560401,
		},
		{
			values:    []float64{-189, 113, 121, -114, -145, 110},
			expectedM: 0.1772088415931945,
		},
	}

	for _, test := range tests {
		m := CalculatePearson(test.values)

		if m != test.expectedM {
			t.Errorf("CalculatePearson(%v) = (%.6f), want (%.6f)", test.values, m, test.expectedM)
		}
	}
}
