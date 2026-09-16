package main
import "fmt"
func main(){
	var names [5]string=[5]string{"hello"}
	prices:=[5]int{10,15,20,30,40}

	fmt.Println(prices)

	 names[2]="world"   // i can set the value  
	fmt.Println(names)  
	new_price:=prices[1:3]
	fmt.Println(new_price) 
	highlight_price:=new_price[1:]
	println(highlight_price)
	fmt.Println(cap(highlight_price),len(highlight_price))

}


 