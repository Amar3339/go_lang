package main

import "fmt"
import "math"

func main() {
	maths()
}
func maths() {
	 annualInvestment,years,exreturn := 1000,10,5.5
	// var exreturn = 5.5
	//  years:= 10
	 result := (float64(annualInvestment) * math.Pow(1+exreturn/100, float64(years)))
	fmt.Println(result)
}
