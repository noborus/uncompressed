// Package uncompressed provides a reader to uncompress the compressed data.
//
// Read the first magic number of the data
// and select the appropriate uncompressed reader.
package uncompressed

import (
	"bytes"
	"compress/bzip2"
	"io"

	"github.com/klauspost/compress/gzip"
	"github.com/klauspost/compress/zstd"
	"github.com/pierrec/lz4"
	"github.com/ulikunitz/xz"
)

// Reader implements uncompressed reader for an io.Reader object.
type Reader struct {
	rd io.Reader
}

// NewReader returns a new uncompressed Reader.
func NewReader(reader io.Reader) *Reader {
	buf := [7]byte{}
	n, err := io.ReadAtLeast(reader, buf[:], len(buf))
	if err != nil {
		if err != io.EOF && err != io.ErrUnexpectedEOF {
			// Should errors not happen?
			return &Reader{
				rd: bytes.NewReader(nil),
			}
		}
	}
	rd := io.MultiReader(bytes.NewReader(buf[:n]), reader)
	var r io.Reader
	switch {
	case bytes.Equal(buf[:3], []byte{0x1f, 0x8b, 0x8}):
		r, err = gzip.NewReader(rd)
	case bytes.Equal(buf[:3], []byte{0x42, 0x5A, 0x68}):
		r = bzip2.NewReader(rd)
	case bytes.Equal(buf[:4], []byte{0x28, 0xb5, 0x2f, 0xfd}):
		r, err = zstd.NewReader(rd)
	case bytes.Equal(buf[:4], []byte{0x04, 0x22, 0x4d, 0x18}):
		r = lz4.NewReader(rd)
	case bytes.Equal(buf[:7], []byte{0xfd, 0x37, 0x7a, 0x58, 0x5a, 0x0, 0x0}):
		r, err = xz.NewReader(rd)
	}
	if err != nil || r == nil {
		r = rd
	}
	return &Reader{
		rd: r,
	}
}

// Read reads data into p.
func (r *Reader) Read(p []byte) (n int, err error) {
	return r.rd.Read(p)
}
