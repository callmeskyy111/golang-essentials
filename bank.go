package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// control structures, loops, switch-cases, writing to Files, error-handling

const accountBalanceFile = "balance.txt" // global-constant

func readBalanceFromFile() (float64, error){
	data, err := os.ReadFile(accountBalanceFile)
	// err - file ("balance.txt") might not be present
	if err != nil{
		// any logic we want
		return 1000.00, errors.New("file doesn't exist ⚠️")
	}
	balanceStr:= string(data)
	balance,err := strconv.ParseFloat(balanceStr,64)
	// err - Not everything can be parsed to a FLOAT (ex: "hello")
	if err != nil{
		// any logic we want
		return 1000.00, errors.New("failed to parse stored balance ⚠️")
	}
	return balance,nil
}

func main(){

var accBalance, err = readBalanceFromFile()
if err !=nil{
	fmt.Println("ERROR:",err)
	fmt.Println("----------------------")
	panic("Exiting the process.. 🔴")
}

fmt.Println("WELCOME to GoBank 🏦!")

// for i:=0; i<2; i++{ ❌ // Not needed here..
// ♾️ loop ☑️
	for{
	fmt.Println("\nYour amount is: $",accBalance)
	fmt.Println("What do you want to do?")
	fmt.Println("1️⃣. Check balance")
	fmt.Println("2️⃣. Deposit")
	fmt.Println("3️⃣. Withdraw")
	fmt.Println("🔘.OTHER - Exit")
	
	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)


// Switch - Alternative to if-else,if,else etc.	
	switch choice{
	case 1:
			fmt.Println("Your balance is: $",accBalance)
	case 2:
		fmt.Print("💰 How much do you wanna deposit?: +$")
		var depositAmt float64 // local scope
		fmt.Scan(&depositAmt)
		if depositAmt <=0{
			fmt.Println("INVALID AMOUNT!.. AMOUNT must not be -ve")
			continue
		}
		accBalance+=float64(depositAmt)
		formattedBalance:= fmt.Sprintf("%.2f",accBalance)
		fmt.Println("Deposited ✅.. Your updated account-balance: $",formattedBalance)
		//writeBalanceToFile(accBalance)
		//! Writing to file ✍🏻📂
		os.WriteFile("balance.txt",[]byte(formattedBalance),0644 ) 
	case 3:
		fmt.Print("💰 How much do you wanna withdraw?: -$")
		var withdrawAmt float64 // local scope
		fmt.Scan(&withdrawAmt)
		if withdrawAmt <=0{
			fmt.Println("INVALID AMOUNT!.. AMOUNT must not be -ve")
			continue
		}
		if withdrawAmt>accBalance{
			fmt.Println("Insufficient Balance :(")
			continue
		}
		accBalance-=float64(withdrawAmt)
		formattedBalance:= fmt.Sprintf("%.2f",accBalance)
		fmt.Println("Amount withdrawn ✅.. Your updated account-balance: $",formattedBalance)
	default:
	fmt.Println("Exiting.. Thanks for choosing GoBank")
	return
	//break
	}
}
}