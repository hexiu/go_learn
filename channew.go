/********************************************************
 Describe :
 Author : jaxiuhe
 Email : axiu@jaxiu.cn
 Createtime : 2018/8/28 上午11:19
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

func Sum(num int64) {
	var count int64
	var i int64
	for i < num {
		count += 1
		if i%100000 == 0 {
			i += 1
			count = 0
		}
	}
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	var num int64
	if len(os.Args) > 1 {
		numi, err := strconv.ParseInt(os.Args[1], 2, 10)
		if err != nil {
			log.Panicln(err)
		}
		num = numi
	} else {
		num = 100000000
	}
	cpuNum := runtime.NumCPU()
	starttime := time.Now()
	for i := 0; i < cpuNum; i++ {
		wg.Add(1)
		go Sum(num)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(starttime).Seconds())
}
