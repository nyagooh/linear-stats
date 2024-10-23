package main

import (
	"bufio"
	"fmt"
	k "linear/linear"
	"os"
	"strconv"
)

func main() {
	if len(os.Args[1:]) == 0 {
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
	m, b := k.LinearRegression(values)
	r := k.CalculatePearson(values)

	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", m, b)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", r)

	defer file.Close()

}
