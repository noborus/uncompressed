# uncompressed

Uncompressed provides an expanded reader from various compressed readers.

Uncompressed judging from the magic number of the first few bytes of the file.

Supported compression formats are gzip, bzip2, zstd, lz4, xz.

## example

```go
    r := uncompressed.NewReader(rr)
    p := make([]byte, 1024)
    r.Read(p)
```

See [_examples/zzcat](_examples/zzcat/main.go) for more details.
