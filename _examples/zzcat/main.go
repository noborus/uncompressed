package main

import (
	"flag"
	"io"
	"os"

	"github.com/noborus/uncompressed"
)

func cat(rr io.Reader) {
	r := uncompressed.NewReader(rr)
	if r == nil {
		panic("nil reader")
	}
	io.Copy(os.Stdout, r)
}

func main() {
	flag.Parse()
	if len(flag.Args()) > 0 {
		for _, arg := range flag.Args() {
			file, err := os.Open(arg)
			if err != nil {
				panic(err)
			}
			cat(file)
			file.Close()
		}
	} else {
		cat(os.Stdin)
	}
}
