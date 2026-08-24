package main 
import "fmt"
var a = 10
func main(){
	var b= 10
	fmt.Print("enter your number ")
	fmt.Scan(&b)
	var result=a*b
	fmt.Print(result)
}
