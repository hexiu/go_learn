package main

import "fmt"

func main() {
	var str = "aaasdfvfdgfv"
	longchar(str)
}

type charIndex struct {
	index int
	len   int
}

// longchar
func longchar(str string) (index string) {
	var Flag = make(map[string]charIndex, 0)
	var i, j = 0, 1
	for ; j < len(str); j++ {
		k := j
		for k > i {
			m := i
			for m < k {
				if str[k] == str[m] {
					Flag[str[i:j]] = charIndex{
						index: i,
						len:   j - i,
					}
					j--
					i++
					break
				}

				m++
			}
			k--
		}
		fmt.Println(i, j)
	}
	Flag[str[i:j]] = charIndex{
		index: i,
		len:   j - i,
	}
	fmt.Println(Flag)
	return
}
