package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	filename := "downlog.gz"
	gzip(filename)
}

func gzip(filename string) (n int) {
	var err error
	gf, err := os.Open(filename)
	if err != nil {
		return
	}
	defer gf.Close()

	// bufsize := 1 << 20
	// bufs := make([]byte, bufsize)
	// buf := bytes.NewBuffer(bufs)
	i := 0
	for {
		i++
		body := make([]byte, 1<<10)
		nr, err := gf.Read(body)
		if err != io.EOF && err != nil {
			log.Panic(err)
		}
		n += nr
		index := bytes.Index(body, []byte{31, 139})
		if index != -1 {
			fmt.Println(i, index)
		}
		if err == io.EOF {
			break
		}
	}
	return
}
