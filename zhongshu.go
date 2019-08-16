package main

import "fmt"

func main() {
	var list = []int{1, 2, 3, 4, 1, 1, 1, 1, 1, 1, 1, 1, 1, 11, 1, 14, 5, 9, 4, 4, 4, 4}
	maj := zhongshu(list)
	fmt.Println(maj)
	return
}

func zhongshu(list []int) (maj int) {
	maj = list[0]
	for i, c := 0, 0; i < len(list); i++ {
		if c == 0 {
			maj = list[i]
			c = 1
		} else {
			if maj == list[i] {
				c++
			} else {
				c--
			}
		}
	}
	return
}
