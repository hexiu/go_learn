package main

import "fmt"

func main() {
	var list = []int{10,10,9,9,9,9,9, 9, 2, 13, 14, 15, 6, 7, 8, 9}
	qsort(list, 0, len(list)-1)
	fmt.Println(list)
	// queryMid(list, 0, len(list)-1)
}

func qsort(list []int, i, j int) {
	mid := qsort1(list, i, j)
	fmt.Println("mid1:", mid)
	if mid == 0 {
		fmt.Println(list)
		return
	}
	qsort(list, i, mid)
	fmt.Println("--------", mid+1, j)
	qsort(list, mid+1, j)
}

func qsort1(list []int, i, j int) (mid int) {
	if j-i < 1 {
		return
	}

	var tmp int
	tmp = list[i]
	var index = i
	for i != j && i < j {
		// for i < j {
		for i < j && tmp <= list[j] {
			j--
		}
		list[i] = list[j]
		fmt.Println("b:", i, j, list)
		for i < j && list[i] <= tmp {
			i++
		}
		list[j] = list[i]
	}
	index = i
	list[index] = tmp
	fmt.Println(i, j, list, tmp, index)
	return i
}

func qsort2(list []int, i, j int) (mid int) {
	// l := i
	// r := j
	// tmp := list[l]
	// for k := 0; k < j; k++ {
	// 	if tmp > list[k] {
			
	// 	}
	// }
	return
}

func queryMid(list []int, i, j int) {
	midd := len(list)/2 - 1
	flag := false
	if len(list)%2 == 0 {
		flag = true
	}

	mid := qsort1(list, i, j)
	fmt.Println("mid:", mid, midd)

	if mid > midd {
		if !flag {
			qsort(list, i, mid)
			fmt.Println("mid:", list[midd])
		} else {
			fmt.Println(list)
			fmt.Println(list[midd-1], list[midd], midd)

			fmt.Println("mid:", float64(list[midd-1]+list[midd])/2.0)
		}
	}
}
