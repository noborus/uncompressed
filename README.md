# uncompressed

Uncompressed provides an expanded reader from various compressed readers.

Uncompressed judging from the magic number of the first few bytes of the file.


## example

```go
    r := uncompressed.NewReader(rr)
    p := make([]byte, 1024)
    r.Read(p)
```

See [_examples/zzcat](blob/master/_examples/zzcat/) for more details.