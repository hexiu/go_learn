/********************************************************
 Describe :
 Author : jaxiuhe
 Email : axiu@jaxiu.cn
 Createtime : 2018/8/21 下午10:59
 CopyRight : Copyright 2018 Tencent Inc.
 Permit : GPL
********************************************************/

package main

import (
	"crypto/sha1"
	"fmt"
	"hash"
)

var count int64 = 1

func main() {
	fmt.Println("hi")
	fmt.Println(hash.Hash())
	fmt.Println(sha1.New())

}
