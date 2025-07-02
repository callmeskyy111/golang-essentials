package main

import (
	"fmt"
	"time"
)

//pipelines
//blocking

// func processNum(numChan chan int){
// for num := range numChan{
// 	fmt.Println("Processing number",num)
// 	time.Sleep(time.Second *1)
// }

// }

//receive
// func sum(result chan int, num1 int, num2 int){
// 	numResult := num1 + num2
// 	result <-numResult
// }

// goRoutine synchronizer
// func task1(done chan bool){
// 	defer func() {done<-true}()
// 	fmt.Println("processing...")
// 	//done<-true
// }

func emailSender(emailChan <-chan string, done chan<- bool){
	defer func() {done <- true}()
	//emailChan<-"something@gmail.com" //❌
	for email := range emailChan {
		fmt.Println("Sending email to..",email)
		time.Sleep(time.Second)
	}
}

func main() {

	chanInt:=make(chan int)
	chanStr:=make(chan string)
	
	go func ()  {
		chanInt <- 10
	}()

	go func ()  {
		chanStr <- "hello"
	}()

	for i:=0; i<2; i++{
		select {
		case chanIntVal:= <-chanInt:
		fmt.Println("Received DATA from chanInt: ",chanIntVal)
		case chanStrVal:= <-chanStr:
			fmt.Println("Received DATA from chanStr: ",chanStrVal)	
		//default:	
		}
	}

	// emailChan := make(chan string, 100)
	// done:=make(chan bool)

	// go emailSender(emailChan, done)

	// for i:=0; i<10;i++{
	// 	emailChan <- fmt.Sprintf("%d@gmail.com",i)
	// }

	// fmt.Println("done sending..")
	// close(emailChan) //closing is important
	// <-done

	// emailChan <- "1@example.com"
	// emailChan <-"2@example.com"
	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)



	// done :=make(chan bool)
	// go task1(done)
	// <- done // block

	// result := make(chan int)

	// go sum(result, 4,5)
	// res := <- result //blocking
	// fmt.Println(res)

	// numChan := make(chan int)

	//  go processNum(numChan)
	// numChan <-5

	// for{
	// 	numChan <- rand.Intn(100)
	// }

	//time.Sleep(time.Second*2)


	// messageChan := make(chan string)
	// messageChan <- "ping!" //sending data
	// msg := <-messageChan //receiving data

	//fmt.Println(msg)
	//fmt.Println()
	
}