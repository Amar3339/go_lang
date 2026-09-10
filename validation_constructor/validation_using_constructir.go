package main

import (
	
	"fmt"
	"strconv"
	"run/reuse"

)

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
// var appUser user
 appUser,err:= reuse.NewUser(firstname,lastname,age,birthDate )
if err!=nil{
	fmt.Print(err)
	return
}

 appUser.OutputUserdetails()
}
// fmt.Println("u.firstname,u.lastname,u.age")


func userData(PromptText string)string{
	fmt.Println(PromptText)
	var value string
	fmt.Scanln(&value)
	
	return value

}
