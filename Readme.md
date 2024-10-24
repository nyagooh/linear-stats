# Linear Regression and Pearson Correlation Coefficient

## Overview

This project implements linear regression and calculates the Pearson correlation coefficient to analyze relationships between data points. These statistical tools are widely used in various fields such as business, finance, healthcare, and social sciences to make predictions and understand data patterns.

## Table of Contents

- [Linear Regression and Pearson Correlation Coefficient](#linear-regression-and-pearson-correlation-coefficient)
  - [Overview](#overview)
  - [Table of Contents](#table-of-contents)
    - [Applications of Linear Regression](#applications-of-linear-regression)
  - [What is the Pearson Correlation Coefficient?](#what-is-the-pearson-correlation-coefficient)
    - [Applications of Pearson Correlation Coefficient](#applications-of-pearson-correlation-coefficient)
  - [Getting Started](#getting-started)
  - [Clone the Repository](#clone-the-repository)
  - [Usage](#usage)
  - [Error Handling](#error-handling)
  - [License](#license)

Linear regression is a statistical method used to find a straight line that best fits a set of data points. It helps us understand the relationship between two variables:

- **Dependent Variable (Y)**: The outcome we want to predict (e.g., house price).
- **Independent Variable (X)**: The predictor that affects the dependent variable (e.g., size of the house).

The formula for simple linear regression is:

$$ Y = mX + b $$

Where:
- **Y** is the predicted value.
- **X** is the independent variable.
- **m** is the slope of the line (how much Y changes for each unit change in X).
- **b** is the y-intercept (the value of Y when X is zero).

### Applications of Linear Regression

Linear regression is used for:
- Making predictions based on existing data.
- Understanding how changes in one variable affect another.

## What is the Pearson Correlation Coefficient?

The Pearson correlation coefficient measures how strongly two variables are related. It gives a value between -1 and 1:

- **1**: Perfect positive correlation (as one variable increases, so does the other).
- **-1**: Perfect negative correlation (as one variable increases, the other decreases).
- **0**: No correlation (no relationship between the variables).

### Applications of Pearson Correlation Coefficient

Correlation is useful for:
- Determining relationships between variables.
- Analyzing trends in data.

## Getting Started

To run this project, ensure you have Go installed on your machine.


## Clone the Repository
 Clone this repository and navigate to the project directory.
 ```bash
 https://github.com/nyagooh/linear-stats.git
 ```
## Usage

To execute the program,Create a txt file where you'll have your data set.:
```bash
touch data.txt
```
Input your data in the file:
```bash
189
113
121
114
145
110
```
Run your program:
```bash
go run . data.txt
```
Expected output:
```bash
Linear Regression Line: y = -8.742857x + 153.857143
Pearson Correlation Coefficient: -0.5330331012
```
## Error Handling
The program runs even in a scenario where the data set has negatives.
```bash
-189
113
121
-114
-145
110
```
The expected output will look something like:
```bash
Linear Regression Line: y = 13.885714 + -52.047619
Pearson Correlation Coefficient: 0.1772088416
```
The program also hanmdles edge cases where the data set has decimal values
```bash
189.8
113.6
121.7
114.9
145.1
110
```
The expected output will look something like this:
```bash
Linear Regression Line: y = -8.894286 + 154.752381
Pearson Correlation Coefficient: -0.5408937201
```
## License