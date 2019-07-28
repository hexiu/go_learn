/********************************************************
 Describe :
 Author : jaxiuhe
 Email : axiu@jaxiu.cn
 Createtime : 2018/8/24 下午4:14
 CopyRight : Copyright 2018 Tencent Inc.
 Permit : GPL
********************************************************/

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.MemStats{})
}
