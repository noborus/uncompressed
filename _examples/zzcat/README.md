# zzcat

zzcat is a command to uncompress and output data.

## Usage

```shell
zcat file
```

You can also extend files of different types.

```shell
ls testdata/test.txt*
```

```console
testdata/test.txt        testdata/test.txt.gz   testdata/test.txt.xz
testdata/test.txt.bzip2  testdata/test.txt.lz4  testdata/test.txt.zstd
```

```shell
zcat testdata/test.*
```

```console
test
test
test
test
test
test
```
