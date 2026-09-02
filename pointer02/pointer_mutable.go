package main


import "fmt"
func main(){
	age:=32
	var agePointer *int  //* shows value of address means value of age=32
	
	agePointer= &age //& it show the memory address of  value by which it get value 
	fmt.Println("age:",*agePointer)
	getadultAge(agePointer)
	fmt.Println(age)

}
func getadultAge(age *int){
// return age-18
*age=*age-18

}