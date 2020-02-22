package uncompressed_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/noborus/uncompressed"
)

func Example() {
	file, err := os.Open(filepath.Join("testdata", "test.txt.zstd"))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	r := uncompressed.NewReader(file)
	fmt.Println(ioutil.ReadAll(r))
}
