package main

import "fmt"

// control structures

func main(){

var accBalance = 1000.0

fmt.Println("WELCOME to GoBank 🏦!")
fmt.Println("Your initial-amount is: $",accBalance)
fmt.Println("What do you want to do?")
fmt.Println("1️⃣. Check balance")
fmt.Println("2️⃣. Deposit")
fmt.Println("3️⃣. Withdraw")
fmt.Println("4️⃣. Exit")

var choice int
fmt.Print("Your choice: ")
fmt.Scan(&choice)

// wantsCheckBalance:= choice ==1

if choice == 1{
fmt.Println("Your balance is: $",accBalance)
} else if choice == 2 {
fmt.Print("💰 How much do you wanna deposit?: +$")
var depositAmt float64 // local scope
fmt.Scan(&depositAmt)
if depositAmt <=0{
	fmt.Println("INVALID AMOUNT!.. AMOUNT must not be -ve")
	return //exits the program
}
accBalance+=float64(depositAmt)
formattedBalance:= fmt.Sprintf("%.2f",accBalance)
fmt.Println("Deposited ✅.. Your updated account-balance: $",formattedBalance)
}else if choice == 3{
	fmt.Print("💰 How much do you wanna withdraw?: -$")
	var withdrawAmt float64 // local scope
    fmt.Scan(&withdrawAmt)
	if withdrawAmt <=0{
		fmt.Println("INVALID AMOUNT!.. AMOUNT must not be -ve")
		return //exits the program
	}
	if withdrawAmt>accBalance{
		fmt.Println("Insufficient Balance :(")
		return
	}
	accBalance-=float64(withdrawAmt)
	formattedBalance:= fmt.Sprintf("%.2f",accBalance)
	fmt.Println("Amount withdrawn ✅.. Your updated account-balance: $",formattedBalance)
}else{
	fmt.Println("Good Day!")
}

fmt.Println("You chose: ",choice)

}