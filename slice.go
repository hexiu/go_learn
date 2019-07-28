/********************************************************
 Describe :
 Author : jaxiuhe
 Email : axiu@jaxiu.cn
 Createtime : 2018/8/17 下午7:15
 CopyRight : Copyright 2018 Tencent Inc.
 Permit : GPL
********************************************************/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("test!")
	slice()
	dict()
	intert()
}

func slice() {
	list := [][]int{[]int{1, 2, 3}, []int{2, 3}, []int{5}}
	fmt.Println(list)
}

func dict() {
	dictt := map[string]map[int]bool{"test": {2: true}, "b": {3: false}}
	fmt.Println(dictt)
	for k, v := range dictt {
		fmt.Println(k, v)
	}
}

func intert() {
	dict := map[interface{}]map[interface{}]interface{}{1: {"haha": "test"}, 2: {2: false}}
	fmt.Println(dict)
	for k, v := range dict {
		switch k.(type) {
		case int:
			fmt.Println(k, v)
			for k1, v1 := range v {
				switch k1.(type) {
				case int:
					fmt.Println("int ", k1, v1)
				case string:
					fmt.Println("string ", k1, v1)

				}
			}
		case string:
			fmt.Println("string:", k)
		}
	}
}
