package main

import "fmt"

//interface
type paymenter interface{
	pay(amount float32)
	refund(amount float32, account string)
}

type payment struct{
	gateway paymenter
}


func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

 type razorPay struct{}
func (r razorPay) pay(amount float32) {
	// logic to make payment
	fmt.Println("💰Making payment using RazorPay..",amount)
}

type stripe struct{}
func (s stripe) pay(amount float32){
	// logic to make payment
	fmt.Println("💸Making payment using STRIPE", amount)
}

type fakePayment struct{}

func (f fakePayment) pay(amount float32){
	fmt.Println("🧪Testing payment using DUMMY GateW..",amount)
}

type payPalPayment struct{}

func(p payPalPayment) pay(amt float32){
	// logic
	fmt.Println("💵Making PAYMENT using PayPal payment-gw..",amt)
}

func (p payPalPayment) refund(amt float32, account string){
	// logic goes here..
}


func main() {
	//stripePaymentGw :=stripe{}
	//razorPayPaymentGw :=razorPay{}
	//fakePaymentGw :=fakePayment{}
	payPalPaymentGw:= payPalPayment{}
	newPayment:=payment{
		gateway: payPalPaymentGw,
	}
	newPayment.makePayment(700)
	
}