package main

import (
	"fmt"
	"math"
)

func main() {
	var annualInvestment = 1000
	var exreturn = 5.5
	var years = 10
	var result = (float64(annualInvestment) * math.Pow(1+exreturn/100, float64(years)))
	fmt.Println(result)
}
