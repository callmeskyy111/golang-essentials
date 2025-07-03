package auth

import "fmt"

//! PACKAGE SCOPE ⚠️
//email -> private
//Email -> Exporting

func LoginWithCred(username string, password string) {
	fmt.Println("Login user using: ", username, password)
}