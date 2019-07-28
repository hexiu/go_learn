/********************************************************
 Describe :
 Author : jaxiuhe
 Email : axiu@jaxiu.cn
 Createtime : 2018/8/19 下午11:46
 CopyRight : Copyright 2018 Tencent Inc.
 Permit : GPL
********************************************************/

package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func hello(s chan string) {
	s <- "hello"
	fmt.Println("ok")
}

func main() {
	numcpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numcpu)
	fmt.Println("this os has cpu num:", numcpu)
	var pronum int
	if len(os.Args) > 2 {
		pron, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Println("Pro num err! ", pron, err)
		} else {
			pronum = pron
		}
	} else {
		pronum = 1
	}
	fmt.Println("this os has pro num:", pronum)

	s := make(chan string)
	go hello(s)
	fmt.Println(<-s)
	//go hello(s)
	//time.Sleep(1*time.Second)
	fmt.Println("done")
	sumall()
}

var su chan int64 = make(chan int64, 1)

var countsum struct {
	count int64
	lock  sync.Mutex
}

func sumcount(i, num int64) {
	length := i + num
	countsum.lock.Lock()
	for j := i; j < length; j++ {
		//countsum++
		countsum.count += j
	}
	countsum.lock.Unlock()
	workgroup.Done()
}

func sum(i, num int64) {
	var s int64 = <-su
	j := i
	length := i + num
	for j < length {
		s += j
		j += 1
	}
	su <- s
	endtime := time.Now()
	fmt.Println("timesub:", s, endtime.Nanosecond(), "pid is :", os.Getpid(), " ppid is ", os.Getppid(), " gid is ", os.Getgid())
	workgroup.Done()
	//fmt.Println("sub time:",midtime.Nanoseconds(),middtime)
}

var workgroup sync.WaitGroup

func sumall() {
	var sumnum int64
	su <- sumnum
	var all int64
	if len(os.Args) >= 2 {
		alln, err := strconv.Atoi(os.Args[3])
		all = int64(alln)
		if err != nil {
			log.Println("all is not num! is :", os.Args[3], err)
		}
	} else {
		all = 10000000
	}
	fmt.Println("all is :", all)

	var slicenum int64
	if len(os.Args) >= 2 {
		slicen, err := strconv.Atoi(os.Args[1])
		slicenum = int64(slicen)
		if err != nil {
			log.Println("args is not num! is :", os.Args[1], err)
		}
	} else {
		slicenum = 10
	}
	fmt.Println("os.Args is :", slicenum)
	var num int64
	num = all / slicenum
	starttime := time.Now()
	var i int64
	for i = 0; i < all/num; i++ {
		workgroup.Add(1)
		go sum(i*num, num)
	}
	workgroup.Wait()
	fmt.Println("sum rel:", <-su, "time is :", time.Now().Sub(starttime).Seconds())

	starttime2 := time.Now()
	//fmt.Println(starttime2)
	for i = 0; i < slicenum; i++ {
		workgroup.Add(1)
		//fmt.Println(i,num)
		go sumcount(i*num, num)
	}
	workgroup.Wait()
	fmt.Println("sum rel2:", slicenum, countsum.count, "time is :", time.Now().Sub(starttime2).Seconds())

	//time.Sleep(1*time.Microsecond)
	var count int64 = 0
	starttime1 := time.Now()
	for i = 0; i < all; i++ {
		count += i
	}
	fmt.Println("count:", count, time.Now().Sub(starttime1).Seconds())

}
