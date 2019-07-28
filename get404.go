/********************************************************
 Describe :
 Author : jaxiuhe
 Email : axiu@jaxiu.cn
 Createtime : 2018/8/22 下午7:58
 CopyRight : Copyright 2018 Tencent Inc.
 Permit : GPL
********************************************************/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

func CheckErr(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

var work sync.WaitGroup

func get404(url string) (error string) {
	work.Add(1)
	clt := http.DefaultClient
	req, err := http.NewRequest("GET", url, nil)
	CheckErr(err)
	resp, err := clt.Do(req)
	fmt.Println(resp.Request.RemoteAddr)
	CheckErr(err)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if strings.Contains(string(body), "not found") {
		work.Done()
		fmt.Println(resp.Request.RequestURI, resp.Request.RemoteAddr)
		return "ok"
	}
	work.Done()
	fmt.Println(resp.Status, string(body), resp.Request)
	return "error"
}

//var url string = "http://10.51.88.133/555555555555/iuthgkhggfjds"
var url string = "http://monitor.djcs.cdntip.com/abcdef"

//var url string = "http://dl_dir.qq.com/abcdef/asdffg"

func main() {
	fmt.Println(get404(url))
	count := 500000
	for i := 0; i < count; i++ {

		go get404(url + strconv.Itoa(i+rand.Int()))
		if i%1000 == 0 {
			fmt.Println("======================: ", i)
			time.Sleep(1 * time.Second)
		}
		if i%3000 == 0 {
			time.Sleep(2 * time.Second)
		}
	}
	work.Wait()
	fmt.Println("done")
}
