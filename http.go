/********************************************************
 Describe :
 Author : jaxiuhe
 Email : axiu@jaxiu.cn
 Createtime : 2018/9/4 下午12:32
 CopyRight : Copyright 2018 Tencent Inc.
 Permit : GPL
********************************************************/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	var url string
	if len(os.Args) > 1 {
		url = os.Args[2]
	} else {
		url = "http://101.226.233.196/img/4021.png"
	}
	httpReq(url)
}

func httpReq(url string) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Host", "img.pet.qq.com")
	clt := http.DefaultClient
	resp, err := clt.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(url)
	fmt.Println(resp)
	fmt.Println(string(body))
}
