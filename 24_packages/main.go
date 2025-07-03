package main

//! PACKAGE SCOPE ⚠️
//email -> private
//Email -> Exporting

import (
	"fmt"

	"github.com/callmeskyy111/golang-package/auth"
	"github.com/callmeskyy111/golang-package/user"
	"github.com/fatih/color" //external package
)

// main - Entry Point
func main(){
	auth.LoginWithCred("skyy111","password12345")
	session:=auth.GetSession()
	fmt.Println(session)

	user:=user.User{
		Email: "user@user.com",
		Name: "User-1",
		Password: "secret1234",
	}

	//fmt.Println(user.Email, user.Name, user.Password)
	color.Red(user.Email)
	color.Green(user.Name)
	color.Cyan(user.Password)
}

