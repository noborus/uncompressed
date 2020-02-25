package uncompressed

import (
	"compress/bzip2"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestReader(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "blank",
			args: args{reader: strings.NewReader("")},
			want: []byte(""),
		},
		{
			name: "err",
			args: args{reader: bzip2.NewReader(strings.NewReader("test"))},
			want: []byte(""),
		},
		{
			name: "plain",
			args: args{reader: strings.NewReader("test")},
			want: []byte("test"),
		},
		{
			name: "long",
			args: args{reader: strings.NewReader("testtesttest")},
			want: []byte("testtesttest"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewReader(tt.args.reader)
			got, err := ioutil.ReadAll(r)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileReader(t *testing.T) {
	testdata := "testdata"
	type args struct {
		filepath string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "plain",
			args: args{filepath: "test.txt"},
			want: []byte("test\n"),
		},
		{
			name: "gz",
			args: args{"test.txt.gz"},
			want: []byte("test\n"),
		},
		{
			name: "bzip2",
			args: args{"test.txt.bzip2"},
			want: []byte("test\n"),
		},
		{
			name: "lz4",
			args: args{"test.txt.lz4"},
			want: []byte("test\n"),
		},
		{
			name: "zstd",
			args: args{"test.txt.zstd"},
			want: []byte("test\n"),
		},
		{
			name: "xz",
			args: args{"test.txt.xz"},
			want: []byte("test\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr, err := os.Open(filepath.Join(testdata, tt.args.filepath))
			if err != nil {
				t.Fatal(err)
			}
			r := NewReader(rr)
			got, err := ioutil.ReadAll(r)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReader() = %v, want %v", got, tt.want)
			}
		})
	}
}
