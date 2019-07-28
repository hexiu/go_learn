package main 

import "fmt"

var sign = make(chan bool,2)
func main()  {
	var a = make(chan int,5)
	go inchan(a)
	go outchan(a)
	<-sign
	<-sign
}

func inchan(ain chan<- int ) {
	for i:=0;i<100;i++{
		ain<-i
	}
	close(ain)
	fmt.Println("Enter all.")
	sign <- true
} 

func outchan(aout <-chan int )  {
	for e:=range aout {
		fmt.Println(e)
		// i := <- e
		// fmt.Println(i)
	}
	fmt.Println("out all.")			

	sign<-true
}