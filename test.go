package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello World!\n")

	list := []int{10, 10, 10, 2, 2, 3, 4, 5, 6, 7, 7, 2, 2, 3, 4, 5, 6, 7, 8, 9}
	hi()
	fmt.Println(list, len(list))
	qsort(list, 0, len(list)-1)
	fmt.Println(list)

	fmt.Println(123>>1, 123<<2)
	return
}

func hi() {
	fmt.Println("hi")
}

func qsort(list []int, lo, hi int) {
	if lo >= hi {
		return
	}
	fmt.Println(list, lo, hi)
	mid := qsortm(list, lo, hi)
	fmt.Println("mid:", mid, lo, hi)
	// return
	qsort(list, lo, mid)
	qsort(list, mid+1, hi)
	return
}

func qsortm(list []int, lo, hi int) (mid int) {
	if lo >= hi {
		return
	}
	l := lo
	r := hi
	key := list[lo]
	for l < r && l != r {

		for l < r && key <= list[r] {
			r--
		}
		list[l] = list[r]
		for l < r && key >= list[l] {
			l++
		}
		list[r] = list[l]
	}
	fmt.Println(l, r, list)
	list[l] = key
	mid = l
	return mid
}
