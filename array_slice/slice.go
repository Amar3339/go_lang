package main
import "fmt"
func main(){
	
	prices:=[]int{10,15,20,30,40}
prices[3]=100
// prices[6]=99 //index out of range [6] with length 5 this error show by acessing or changing array which is not in slice 
	new_price:=append(prices,25)  //it makes size 6 and brand new underlying array ,with capcity increases this is like dynamic_array_type  
fmt.Println(new_price,prices) //original slice didn't change by go .if you want to change you have to reasign prices in the place of new_price 



	 

}


 