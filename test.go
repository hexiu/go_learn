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


	nums := make([]int,0)
	for i:=1;i<=100;i++ {
		nums = append(nums,i)
	}
	fmt.Println(nums)
	num:=0
	for i:=0;num!=5050;{
		if (i+1)%11 == 0{
			if num[i] != 0 {

			}else{
				
			}
		}
		if nums[i]== 0{
			num++
		}
		if i==99 {
			i=0
		}else{
			i++
		}
	}
	fmt.Println(num)
	fmt.Println(nums)


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
