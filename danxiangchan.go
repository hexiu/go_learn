package main 

import "fmt"

var achan = make(chan int, 10)
var sign = make(chan int ,2)

func main()  {
	go InputChan(achan)
	go OutputChan(achan)
	<-sign
	<-sign
	go InputChan(achan)
	go OutputChan(achan)
	<-sign
	<-sign
	go InputChan(achan)
	go OutputChan(achan)
	<-sign
	<-sign
}

func InputChan(ain chan<- int)  {
	for i :=0;i<10;i++{
		fmt.Println(i)
		ain <- i
	}
	fmt.Println("Enter all.")
	sign<-0
}

func OutputChan(aout <-chan int)  {
	for i :=0;i<10;i++{
		fmt.Println(<-aout)
	}
	fmt.Println("Output all.")
	sign <-1
}