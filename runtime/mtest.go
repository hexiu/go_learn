package main

import "fmt"
import "runtime"

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.Version())
	
}
