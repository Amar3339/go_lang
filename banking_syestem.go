// normal banking syestem in which take input from user deposit ,user choice .add balance ,debit balance ,update balance

package main

import "fmt"

func main() {
	fmt.Println("welcome to our Go Bank")

	var balance = 10000.0
	for {
		fmt.Print("what do you want to do ")
		fmt.Println("1. : chek balance  ")
		fmt.Println("2. : deposit money")
		fmt.Println("3. : withdraw money  ")
		fmt.Println("4. : exit")

		var choice int
		fmt.Println("your choice ")
		fmt.Scan(&choice)
		if choice == 1 {
			fmt.Println(" your balance is ", balance)
		} else if choice == 2 {
			fmt.Print("enter your deposit ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("deposit should be valid amount")
				return
			}
			balance += depositAmount //   balance=balance+depositAmmount
			fmt.Println(balance)
		} else if choice == 3 {
			fmt.Println("enter your withdraw amount  ")
			var withdrawAmount float64
			fmt.Scanf("%f", &withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("please enter valid ammount ")
				return
			} else if withdrawAmount > balance {
				fmt.Println("please enter valid ammount ")
				return
			}
			balance -= withdrawAmount

			fmt.Println("your total balance", balance)
		} else if choice == 4 {
			fmt.Println("bye")
			break

		}

		fmt.Println(choice)
	}
fmt.Println("thanks for banking with us 4")
}