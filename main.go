package main

import (
	"bufio"
	"fmt"
	k "linear/linear"
	"os"
	"strconv"
)
// main is the entry point for the linear regression and correlation coefficient calculation program.//+
// It reads a file provided as a command-line argument, calculates the linear regression line and//+
// Pearson correlation coefficient, and prints the results.//+
func main() {
	if len(os.Args[1:]) == 0 ||  len(os.Args) >2{
		return
	}
	args := os.Args[1:][0]
	file, err := os.Open(args)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	scanner := bufio.NewScanner(file)
	var values []float64
	for scanner.Scan() {
		line := scanner.Text()
		value, _ := strconv.ParseFloat(line, 64)
		values = append(values, value)
	}
	if len(values) <=1 {
		fmt.Println("Not enough data points to perform linear regression and  Correlation Coefficient.")
        return
	}
	m, b := k.LinearRegression(values)
	r := k.CalculatePearson(values)

	fmt.Printf("Linear Regression Line: y = %.6f + %.6f\n", m, b)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", r)

	defer file.Close()

}
