# uncompressed Reader

![Go](https://github.com/noborus/uncompressed/workflows/Go/badge.svg)

## No need to import. Reference implementation.

Uncompressed provides uncompressed reader from various compressed readers.

Uncompressed reader identifies the file from the magic number in the first few bytes.

Supported compression formats are gzip, bzip2, zstd, lz4, xz.

## example

```go
package main

import (
	"io"
	"os"

	"github.com/noborus/uncompressed"
)

func main() {
	file, err := os.Open("test.txt.zstd")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	r := uncompressed.NewReader(file)
	io.Copy(os.Stdout, r)
}
```

See [_examples/zzcat](_examples/zzcat/main.go) for more details.
