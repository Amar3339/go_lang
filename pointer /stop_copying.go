package main
import "fmt"
func main(){
	age:=32
	var agePointer *int  //* shows value of address means value of age=32
	
	agePointer= &age //& it show the memory address of  value by which it get value 
	fmt.Println("age:",*agePointer)
	adultAge:=getadultAge(*agePointer)
	fmt.Println(adultAge)

}
func getadultAge(age int) int{
return age-18
}