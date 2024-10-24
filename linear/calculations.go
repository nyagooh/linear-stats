package linear

import "math"

//y=mx+b
// LinearRegression calculates the linear regression line coefficients (slope and y-intercept)
// for a given set of y-values. The x-values are assumed to be the indices of the y-values.
//
// The function takes a slice of float64 values representing the y-values.
//
// It returns two float64 values:
// - The first value is the slope (m) of the regression line.
// - The second value is the y-intercept (b) of the regression line.
func LinearRegression(values []float64) (float64, float64) {
	n := float64(len(values))
	var sumX, sumY, sumXY, sumX2 float64

	for i, y := range values {
		x := float64(i)
		sumX += x
		sumY += y
		sumXY += x * y
		sumX2 += x * x
	}
	// m = N(∑x2) - (∑x)² / N(∑xy) - (∑x)(∑y)
	m := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	// b = N(∑y) - m(∑x) / N
	b := (sumY - m*sumX) / n

	return m, b
}

// CalculatePearson calculates the Pearson correlation coefficient between a set of y-values.//+
// The x-values are assumed to be the indices of the y-values.//+
//
// The Pearson correlation coefficient measures the linear relationship between two datasets.//+
// It ranges from -1 to 1, where -1 indicates a strong negative relationship, 0 indicates no relationship,//+
// and 1 indicates a strong positive relationship.//+
////+
func CalculatePearson(values []float64) float64 {
	n := float64(len(values))
	var sumX, sumY, sumXY, sumX2, sumY2 float64

	for i, y := range values {
		x := float64(i)
		sumX += x
		sumY += y
		sumXY += x * y
		sumX2 += x * x
		sumY2 += y * y
	}
	//r=[N∑x2−(∑x)2][N∑y2−(∑y)2]
	// ​N(∑xy)−(∑x)(∑y)​
	numerator := n*sumXY - sumX*sumY
	denominator := math.Sqrt((n*sumX2 - sumX*sumX) * (n*sumY2 - sumY*sumY))

	return numerator / denominator
}
