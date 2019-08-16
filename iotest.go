package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/hexiu/utils/base"
)

func main() {
	iotest()
}

func iotest() {
	argvs := os.Args
	var fp string
	var num int
	if len(argvs) > 2 {
		fp = argvs[1]
		num, _ = strconv.Atoi(argvs[2])
	} else {
		fp = os.TempDir()
		num = 1
	}

	dirs, err := ioutil.ReadDir(fp)
	if err != nil {
		panic(err)
	}
	// fmt.Println(dirs)
	fmt.Println(fp)
	// printDir(dirs)
	t, err := base.NewTicket(num)
	if err != nil {
		panic(err)
	}
	readTest(fp, dirs, t)
}

func printDir(dirs []os.FileInfo) {
	for _, f := range dirs {
		fmt.Println(f)
	}
	return
}

func readTest(fp string, dirs []os.FileInfo, tickets base.Ticket) {
	for _, f := range dirs {
		fmt.Println("Filename:", f.Name())
		if f.IsDir() {
			continue
		}
		filep := filepath.Join(fp, f.Name())
		err := tickets.Take()
		if err != nil {
			return
		}
		fmt.Println("ReadFile:", filep)
		go readFile(filep, tickets)
	}

}

func readFile(fpname string, tickets base.Ticket) {
	defer tickets.Return()

	f, err := os.Open(fpname)
	if err != nil {
		panic(err)
	}

	body := make([]byte, 0)
	var n int
	for err != io.EOF || err == nil {
		var num int
		num, err = f.Read(body)
		n += num
	}
	fmt.Println(fpname, n)
	return
}
