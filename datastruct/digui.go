package main

import "fmt"

func main() {
	var num = 10000
	sum1(num)
	sum2(num)
	a := "hahaa123456"
	fmt.Println(strings(a, 0, len(a)-1))
	fmt.Println(sum3(0, 100))
	var x1, x2 int
	var A []int = []int{1, 4, 2, 1, 3, 5, 7, 3, 9, 2, 1, 7, 3, 5, 6}
	getmaxlow(A, 0, len(A)-1, &x1, &x2)
	fmt.Println(x1, x2)
}

func sum1(num int) {
	var sum int
	for i := 0; i <= num; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func sum2(num int) {
	sum := diedai(num)
	fmt.Println(sum)
}

func diedai(n int) int {
	if n == 1 {
		return 1
	}
	return diedai(n-1) + n
}

func strings(a string, i, j int) string {
	if i < j {
		a = swap(a, i, j)
		fmt.Println(a)
		return strings(a, i+1, j-1)
	} else {
		return a
	}

}

func swap(a string, i, j int) string {
	// fmt.Println(a[0])
	tn := []byte(a)
	tm := tn[i]
	tn[i] = tn[j]
	tn[j] = tm
	return string(tn)
}

func sum3(num1, num2 int) int {
	var mid int
	mid = (num1 + num2) >> 1
	if num1 == num2 {
		return mid
	}

	return sum3(num1, mid) + sum3(mid+1, num2)
}

func getmaxlow(a []int, lo, hi int, x1, x2 *int) {
	if lo+2 == hi || lo+3 == hi {
		var x1L, x2L, x1R, x2R int
		if a[lo] < a[lo+1] {
			x1L = a[lo+1]
			x2L = a[lo]
		} else {
			x1L = a[lo]
			x2L = a[lo+1]
		}
		if lo+2 == hi {
			x1R = a[hi]
		} else {
			if a[lo+2] < a[lo+3] {
				x2R = a[lo+2]
				x1R = a[lo+3]

			} else {
				x1R = a[lo+2]
				x2R = a[lo+3]
			}
		}
		if x1L > x1R {
			*x1 = x1L
			if x2L < x1R {
				*x2 = x1R
			} else {
				*x2 = x2L
			}
		} else {
			*x1 = x1R
			if x1L < x2R {
				*x2 = x2R
			} else {
				*x2 = x1L
			}
		}
		fmt.Println("A:", a[lo:hi+1])
		fmt.Println("x: ", x1L, x2L, x1R, x2R)
		fmt.Println("last:", lo, hi, *x1, *x2)
		return
	}
	mid := (lo + hi) >> 1
	var x1L, x2L, x1R, x2R int
	getmaxlow(a, lo, mid, &x1L, &x2L)
	getmaxlow(a, mid+1, hi, &x1R, &x2R)
	if x1L > x1R {
		*x1 = x1L
		if x2L < x1R {
			*x2 = x1R
		} else {
			*x2 = x2L
		}
	} else {
		*x1 = x1R
		if x1L < x2R {
			*x2 = x2R
		} else {
			*x2 = x1L
		}
	}
	fmt.Println(lo, hi, *x1, *x2)
}
