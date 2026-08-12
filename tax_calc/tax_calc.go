// Build a profit calculator ,ask for Revenie,expense,TaxRate.Calculate earning bfore tax earning after tax
// calculator ration(Ebt/Profit) 
// output Ebt,profit and the ratioh


package main
import "fmt"
func main(){
	var Revenue float64
	var Expense float64
	var taxrate float64
	fmt.Print("enter your revenue")
	fmt.Scan(&Revenue)
	fmt.Print("enter your Expense")
	fmt.Scan(&Expense)
	fmt.Print("enter your taxrate")
	fmt.Scan(&taxrate)
	  var Ebt= Revenue-Expense
 var Profit =Ebt*(1- taxrate/100)
	var ratio= Ebt/Profit
	fmt.Println(Ebt,Profit,ratio)


}
