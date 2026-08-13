package main

import (
	"fmt"
	"math"
)

const inflaionRate = 2.5

func main() {
	var investment float64
	var years float64
	var expectationReturnRate float64 = 5.5
	output("please enter your investment amount \n")
	//  fmt.Print("please enter your investment amount:  ")
	fmt.Scan(&investment)
	output("please enter years")
	//  fmt.Print("please enter years:   ")
	fmt.Scan(&years)
	futurevalue, futurerealvalue := calculatefv(investment, expectationReturnRate, years)
	//  futurevalue:= investment*math.Pow(1+expectationReturnRate/100,years)
	//  futurerealvalue:=futurevalue/math.Pow(1+inflaionRate/100,years)
	formattedFV := fmt.Sprintf("Fututre value: %.1f\n", futurevalue)
	formattedRFV := fmt.Sprintf("Fututre value: %.1f\n", futurerealvalue)
	// fmt.Println("future value:",futurevalue)
	fmt.Println(formattedFV, formattedRFV)

	// fmt.Printf("future real value %.5f\n",futurerealvalue)
}
func output(text string) {
	fmt.Print(text)

}

func calculatefv(investment, expectationReturnRate, years float64) (float64, float64) {

	FV := investment * math.Pow(1+expectationReturnRate/100, years)
	RFV := FV / math.Pow(1+inflaionRate/100, years)
	return FV, RFV

}
