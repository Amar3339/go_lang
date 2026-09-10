package reuse
import (
	"errors"
	"fmt"
	
	"time"
)
type User struct{
	firstName string
	lastname string
	age int
	birthDate string
	createdAt time.Time

}
func (u User ) OutputUserdetails(){  
	fmt.Println(u.firstName,u.lastname,u.birthDate,u.age)
}
func NewUser(firstname string,lastname string,age int  ,birthDate string  ) (*User ,error) {
	if firstname=="" || lastname==""{
		return nil,errors.New("fisrtname ,lastname should be required ")
	}


return &User {


	firstName: firstname,
	lastname: lastname,
	age: age,
	birthDate: birthDate,
	createdAt:time.Now(),
},nil
}