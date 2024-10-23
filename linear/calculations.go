package linear

import "math"

//y=mx+b
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
	//m=N(∑x2)−(∑x)2N(∑xy)−(∑x)(∑y)​
	m := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	//b=N∑y−m(∑x)​
	b := (sumY - m*sumX) / n

	return m, b
}

//how things are related.
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
