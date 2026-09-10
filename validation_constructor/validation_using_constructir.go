package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
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
// var appUser user
appUser,err:= newuser(firstname,lastname,age,birthDate )
if err!=nil{
	fmt.Print(err)
	return
}

 appUser.outputUserdetails()
}
// fmt.Println("u.firstname,u.lastname,u.age")


func userData(PromptText string)string{
	fmt.Println(PromptText)
	var value string
	fmt.Scanln(&value)
	
	return value

}
func (u user ) outputUserdetails(){  
	fmt.Println(u.firstName,u.lastname,u.birthDate,u.age)
}
func newuser(firstname string,lastname string,age int  ,birthDate string  ) (*user ,error) {
	if firstname=="" || lastname==""{
		return nil,errors.New("fisrtname ,lastname should be required ")
	}


return &user {


	firstName: firstname,
	lastname: lastname,
	age: age,
	birthDate: birthDate,
	createdAt:time.Now(),
},nil
}