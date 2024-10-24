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
			values:    []float64{189,113,121,114,145,110},
			expectedM: -8.742857,
			expectedB: 153.857143,
		},
		{
			values:    []float64{189.8,113.6,121.7,114.9,145.1,110},
			expectedM:  -8.894286,
			expectedB: 154.752381,
		},
		{
			values:    []float64{-189,113,121,-114,-145,110},
			expectedM: 13.885714,
			expectedB: -52.047619,
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
			values:    []float64{189.8,113.6,121.7,114.9,145.1,110},
			expectedM:  -0.5408937201,
		},
		{
			values:    []float64{189,113,121,114,145,110},
			expectedM: -0.5330331012,
			
		},
		{
			values:    []float64{-189,113,121,-114,-145,110},
			expectedM:  0.1772088416,
		},
	}

	for _, test := range tests {
		m := CalculatePearson(test.values)

		if m != test.expectedM  {
			t.Errorf("CalculatePearson(%v) = (%.6f), want (%.6f)", test.values, m,test.expectedM)
		}
	}
}
