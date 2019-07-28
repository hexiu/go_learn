package main

import (
	"fmt"

	"github.com/hexiu/utils/datastruct"
)

func main() {
	vector := datastruct.NewVector()
	vector.Insert(0, 1)
	vector.Insert(0, 1)
	vector.Insert(0, 1)
	vector.Insert(0, 1)
	vector.Insert(0, 1)
	vector.Insert(0, 1)
	vector.Insert(0, 1)
	vector.Insert(0, 1)
	vector.Insert(0, 1)
	vector.Insert(0, 1)
	vector.Put(3, 2)
	fmt.Println(vector.Get(4))
}
