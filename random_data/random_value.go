package main

// package main

import "fmt"
 import "os"
 import "strconv"
 import "github.com/Pallinder/go-randomdata"
 const acc_balance="balance.txt" //this is files (balance.txt) i created and store the value of balance 
 func getbalanceFromfiles()float64{
	data,_:=os.ReadFile(acc_balance)
	balanceText:=string(data)
	balance,_:=strconv.ParseFloat(balanceText,64)
	return balance
 }


 func WriteBalance(balance float64){
 balanceText:=fmt.Sprint(balance)
 os.WriteFile("balance.txt",[]byte(balanceText),0644)
 }


func main() {

    fmt.Println("Welcome to our Go Bank")

    var balance = getbalanceFromfiles()
	genratecustomer()


    for {
		combine()
        
        var choice int

        fmt.Print("Your choice: ")
        fmt.Scan(&choice)

        switch choice {

        case 1:
            fmt.Println("Your balance is:", balance)

        case 2:
            fmt.Print("Enter your deposit: ")

            var depositAmount float64
            fmt.Scan(&depositAmount)

            if depositAmount <= 0 {
                fmt.Println("Deposit should be a valid amount")
                continue
            }

            balance += depositAmount

            fmt.Println("Your total balance:", balance)
            WriteBalance(balance)

        case 3:
            fmt.Print("Enter your withdrawal amount: ")

            var withdrawAmount float64
            fmt.Scan(&withdrawAmount)

            if withdrawAmount <= 0 || withdrawAmount > balance {
                fmt.Println("Please enter a valid amount")
                continue
            }

            balance -= withdrawAmount

            fmt.Println("Your total balance:", balance)
WriteBalance(balance)
        case 4:
            fmt.Println("Bye")
            return

        default:
            fmt.Println("Invalid choice")
            return
        }

        fmt.Println("You selected:", choice)
    }

}
func combine(){
	fmt.Println("What do you want to do?")
        fmt.Println("1. Check balance")
        fmt.Println("2. Deposit money")
        fmt.Println("3. Withdraw money")
        fmt.Println("4. Exit")

}
func genratecustomer(){
	fmt.Println("you can reach us 24*7",randomdata.PhoneNumber())
	fmt.Println("firstname is",randomdata.FirstName(randomdata.Male ))
	fmt.Println("lastname is",randomdata.LastName())
	fmt.Println("addres:=",randomdata.Address())
	fmt.Println("fullname:",randomdata.FullName(0))
	fmt.Println("email:",randomdata.Email())
	fmt.Println("country",randomdata.Country(randomdata.FullCountry))
	fmt.Println("city",randomdata.City())
	fmt.Println("countrycode",randomdata.Country(randomdata.ThreeCharCountry))
	fmt.Println("",randomdata.Paragraph())
	fmt.Println("IPv4:", randomdata.IpV4Address())
}