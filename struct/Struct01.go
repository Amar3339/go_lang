package main

import (
	"fmt"
	"time"
	"strconv"
)
type user struct{
	firstName string
	lastname string
	age int
	birthDate string
	createdAt time.Time

}
func main(){
firstname:= userData("enter your firstname: ")
lastname:= userData("enter your lastname: ")
ageString:= userData("enteryour age")
age, err := strconv.Atoi(ageString)

	if err != nil {
		fmt.Println("Please enter a valid age")
		return
	}
birthDate:= userData("enter your birthDate: ")

// fmt.Print(firstname,lastname,age, birthDate)
var appUser user
appUser=user{
	firstName: firstname,
	lastname: lastname,
	age: age,
	birthDate: birthDate,
	createdAt:time.Now(),
}
 outputUserdetails(&appUser) 
}
// fmt.Println("u.firstname,u.lastname,u.age")


func userData(PromptText string)string{
	fmt.Println(PromptText)
	var value string
	fmt.Scan(&value)
	
	return value

}
func outputUserdetails(u *user){
	fmt.Println(u.firstName,u.lastname,u.birthDate,u.age)
}