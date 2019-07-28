package main 

import "fmt"
import "time"

var sign = make(chan bool,2)
var timeout = make(chan bool,1)

func main()  {
	// test()
	timeOuttest()
}
func test()  {
	num := 100000
	var chan1000 = make(chan int,2)
	go func () {
		time.Sleep(time.Millisecond)
		timeout<-false
	}()
	go func ()  {
		for i:=0;i<num;i++ {
			chan1000 <-i
			fmt.Println("i:",i)
		}
		close(chan1000)
		sign <- true
	}()

	go func ()  {
		var e int 
		var ok bool 
		ok = true
		for { 
		select {
			case e,ok=<-chan1000:
			if !ok {
				fmt.Println("end.")
				break
			}else{
				fmt.Println(e)
			}	
			case ok=<-timeout:
				fmt.Println("timeout.")
				sign<-false
				break
			}
		
			if !ok {
				sign <- true
				break
			}
		}	
		}()
	
<-sign
<-sign
	
	fmt.Println("over")
}


func timeOuttest(){
	num := 10
	var chan1000 = make(chan int,2)
	
	go func ()  {
		for i:=0;i<num;i++ {
			chan1000 <-i
			fmt.Println("i:",i)
			if i % 5 ==0 {
				time.Sleep(2*time.Millisecond)
			}
		}
		close(chan1000)
		sign <- true
	}()

	go func ()  {
		for {
			ok := true
			goon := true
			var e int 
			select {
			case e,ok = <-chan1000:
				if !ok {
					fmt.Println("end")
					break
				}
				fmt.Println(e)
			case goon = <- func() (timeout chan bool){
				timeout = make(chan bool,1)
				go func () {
					time.Sleep(time.Millisecond)
					timeout<-false
				}()
				return timeout
			}() :  
				fmt.Println("timeout")
				break
			}
			if !goon {
				fmt.Println("goon:",goon,ok)
				continue
			}
			
			if !ok {
				fmt.Println("break")
				break
			}
		}	
		sign <-true
	}()
	<-sign 
	<-sign
}