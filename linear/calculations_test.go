package linear

import "testing"

func TestLinearRegression(t *testing.T) {
tests := []struct {
	values      []float64
	expectedM  float64
	expectedB  float64
}{
	{
		values:     []float64{1, 2, 3, 4, 5},
		expectedM:  1.0,
		expectedB:  1.0,
	},
	{
		values:     []float64{2, 4, 6, 8, 10},
		expectedM:  2.0,
		expectedB:  0.0,
	},
	{
		values:     []float64{5, 4, 3, 2, 1},
		expectedM:  -1.0,
		expectedB:  5.0,
	},
	{
		values:     []float64{0, 0, 0},
		expectedM:  0.0,
		expectedB:  0.0,
	},
}

for _, test := range tests {
	m, b := LinearRegression(test.values)
	
	if m != test.expectedM || b != test.expectedB {
        t.Errorf("LinearRegression(%v) = (%.2f, %.2f), want (%.2f, %.2f)", test.values, m, b, test.expectedM, test.expectedB)
    }
}
}

